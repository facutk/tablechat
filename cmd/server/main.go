package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/facutk/tablechat/internal/database"
	"github.com/facutk/tablechat/internal/db"
	"github.com/facutk/tablechat/views"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type App struct {
	queries *db.Queries
	ctx     context.Context
}

// registerMessageRoutes attaches /message endpoints to the given chi router.
func (app *App) registerMessageRoutes(r chi.Router) {
	r.Get("/message", app.messagePageHandler)         // full page
	r.Get("/message/fragment", app.messageGetHandler) // HTMX-only fragment (optional)
	r.Post("/message", app.messagePostHandler)        // receives updates via htmx
}

func (app *App) readMessage() (string, error) {
	// Use ID 1 for the message, or adjust as needed
	msg, err := app.queries.GetMessage(app.ctx, 1)
	if err != nil {
		// Handle the case where message doesn't exist yet
		// You might want to return an empty string or create a default message
		return "", nil // or handle differently based on your needs
	}
	return msg.Message, nil
}

func (app *App) messagePageHandler(w http.ResponseWriter, r *http.Request) {
	// render a simple page and inject the templ fragment
	msg, err := app.readMessage()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	templ.Handler(views.Message(msg)).ServeHTTP(w, r)
}

// returns only the fragment (useful if you want to fetch it directly)
func (app *App) messageGetHandler(w http.ResponseWriter, r *http.Request) {
	msg, err := app.readMessage()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	templ.Handler(views.Message(msg)).ServeHTTP(w, r)
}

func (app *App) messagePostHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	m := r.PostFormValue("message")

	_, err := app.queries.UpdateMessage(app.ctx, db.UpdateMessageParams{
		ID:      1, // Assuming you're updating the message with ID 1
		Message: m,
	})
	if err != nil {
		http.Error(w, "could not save message to database", http.StatusInternalServerError)
		return
	}

	// return updated fragment so HTMX can swap it in-place
	templ.Handler(views.Message(m)).ServeHTTP(w, r)
}

func main() {
	_ = godotenv.Load()

	// Initialize database connection
	cfg := database.DefaultConfig()
	dbConn, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbConn.Close()

	// Run migrations
	if err := database.RunMigrations(dbConn); err != nil {
		log.Fatal("Migrations failed:", err)
	}

	// Initialize sqlc queries
	queries := db.New(dbConn)
	ctx := context.Background()

	// Create app instance with dependencies
	app := &App{
		queries: queries,
		ctx:     ctx,
	}

	// message, err := queries.CreateMessage(ctx, "Hello, World!")
	// if err != nil {
	// 	log.Printf("Note: Could not create test message: %v", err)
	// } else {
	// 	log.Printf("Created test message with ID: %d", message.ID)
	// }

	// Example: Get the message back
	retrievedMessage, err := queries.GetMessage(ctx, 1)
	if err != nil {
		log.Printf("Could not retrieve message: %v", err)
	} else {
		log.Printf("Retrieved message: %s", retrievedMessage.Message)
	}

	gitsha, ok := os.LookupEnv("GITSHA")
	if !ok || gitsha == "" {
		gitsha = "unknown"
	}
	component := views.Hello(fmt.Sprintf("tablechat.me (%s)", gitsha))

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r2 *http.Request) {
		templ.Handler(component).ServeHTTP(w, r2)
	})
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	app.registerMessageRoutes(r)

	// server static files
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/static/", fs)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", r)
}
