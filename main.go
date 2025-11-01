package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

func main() {
	gitsha, ok := os.LookupEnv("GITSHA")
	if !ok || gitsha == "" {
		gitsha = "unknown"
	}
	component := hello(fmt.Sprintf("tablechat.me (%s)", gitsha))
	http.Handle("/", templ.Handler(component))

	// server static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", fs)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
