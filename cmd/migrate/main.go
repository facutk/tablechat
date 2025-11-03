package main

import (
	"log"

	"github.com/facutk/tablechat/internal/database"
)

func main() {
	cfg := database.Config{
		DBPath: "./app.db", // or from environment variable
	}

	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	log.Println("Migrations completed successfully")
}
