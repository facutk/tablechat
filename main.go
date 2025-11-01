package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const messagePath = "tmp/message.txt"

// registerMessageRoutes attaches /message endpoints to the given chi router.
func registerMessageRoutes(r chi.Router) {
	r.Get("/message", messagePageHandler)         // full page
	r.Get("/message/fragment", messageGetHandler) // HTMX-only fragment (optional)
	r.Post("/message", messagePostHandler)        // receives updates via htmx
}

func readMessage() (string, error) {
	b, err := os.ReadFile(messagePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	return string(b), nil
}

func messagePageHandler(w http.ResponseWriter, r *http.Request) {
	// render a simple page and inject the templ fragment
	msg, err := readMessage()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	templ.Handler(message(msg)).ServeHTTP(w, r)
}

// returns only the fragment (useful if you want to fetch it directly)
func messageGetHandler(w http.ResponseWriter, r *http.Request) {
	msg, err := readMessage()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	templ.Handler(message(msg)).ServeHTTP(w, r)
}

func messagePostHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	m := r.PostFormValue("message")

	if err := os.MkdirAll(filepath.Dir(messagePath), 0o755); err != nil {
		http.Error(w, "could not create data dir", http.StatusInternalServerError)
		return
	}
	if err := os.WriteFile(messagePath, []byte(m), 0o644); err != nil {
		http.Error(w, "could not save message", http.StatusInternalServerError)
		return
	}

	// return updated fragment so HTMX can swap it in-place
	templ.Handler(message(m)).ServeHTTP(w, r)
}

func main() {
	_ = godotenv.Load()

	gitsha, ok := os.LookupEnv("GITSHA")
	if !ok || gitsha == "" {
		gitsha = "unknown"
	}
	component := hello(fmt.Sprintf("tablechat.me (%s)", gitsha))

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r2 *http.Request) {
		templ.Handler(component).ServeHTTP(w, r2)
	})
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	registerMessageRoutes(r)

	// server static files
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/static/", fs)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", r)
}
