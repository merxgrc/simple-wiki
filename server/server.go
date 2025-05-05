package main

import (
	"fmt"
	"log"
	"net/http"
)

// This function handles http requests.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[0:])
}

// Main function sets up an http server that listens on port 8080 and uses the handler function for all routes.
func main() {
	http.HandleFunc("/", handler)                // Registers the handler function when a request matches the "/" pattern
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start an http web server on port 8080, crash if error.
}
