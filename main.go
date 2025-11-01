package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

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

	// server static files
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/static/", fs)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", r)
}
