package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"

	"github.com/ccadden/htmx-go/web_1.0/views"
)

func main() {
	component := views.Hello("World")
	http.Handle("/", templ.Handler(component))

	http.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Contacts")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
