package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("tablechat.me")
	http.Handle("/", templ.Handler(component))

	// server static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", fs)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
