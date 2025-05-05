package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[0:])
}

// Main function sets up an http server that listens on port 8080 and uses the handler function for all routes.
func main() {
	http.HandleFunc("/", handler) // handler is just the name I decided to give the function it can be anything.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
