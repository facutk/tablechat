package database

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "modernc.org/sqlite" // Pure Go SQLite driver
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type Config struct {
	DBPath string
}

func DefaultConfig() Config {
	return Config{
		DBPath: "./data/app.db",
	}
}

func NewConnection(cfg Config) (*sql.DB, error) {
	// Create directory if it doesn't exist
	dbDir := "./data"
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create data directory: %w", err)
		}
		log.Printf("Created database directory: %s", dbDir)
	}

	// Create database file if it doesn't exist
	if _, err := os.Stat(cfg.DBPath); os.IsNotExist(err) {
		file, err := os.Create(cfg.DBPath)
		if err != nil {
			return nil, fmt.Errorf("failed to create database file: %w", err)
		}
		file.Close()
		log.Printf("Created database file: %s", cfg.DBPath)
	}

	db, err := sql.Open("sqlite", cfg.DBPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	log.Printf("Database connection established: %s", cfg.DBPath)
	return db, nil
}

func RunMigrations(db *sql.DB) error {
	// List all embedded migration files for debugging
	log.Println("Checking embedded migration files...")
	err := fs.WalkDir(migrationsFS, "migrations", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			log.Printf("Found migration file: %s", path)
		}
		return nil
	})
	if err != nil {
		log.Printf("Error walking migration files: %v", err)
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	source, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return fmt.Errorf("failed to create migration source: %w", err)
	}

	m, err := migrate.NewWithInstance("iofs", source, "sqlite3", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	log.Println("Running database migrations...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return fmt.Errorf("failed to get migration version: %w", err)
	}

	log.Printf("Database migrated to version: %d (dirty: %t)", version, dirty)

	// Verify the messages table was created
	var messagesTableExists bool
	err = db.QueryRow(`
		SELECT COUNT(*) > 0 
		FROM sqlite_master 
		WHERE type='table' AND name='messages'
	`).Scan(&messagesTableExists)

	if err != nil {
		log.Printf("Error checking for messages table: %v", err)
	} else if messagesTableExists {
		log.Printf("Messages table exists: %t", messagesTableExists)
	} else {
		log.Printf("WARNING: Messages table does not exist after migration!")
		// Let's try to create it directly as a fallback
		log.Println("Attempting to create messages table directly...")
		_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS messages (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				message TEXT NOT NULL,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP
			)
		`)
		if err != nil {
			log.Printf("Failed to create messages table directly: %v", err)
		} else {
			log.Println("Successfully created messages table directly")
		}
	}

	return nil
}
