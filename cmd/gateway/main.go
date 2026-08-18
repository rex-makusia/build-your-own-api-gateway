package main

import (
	"fmt"
	"log"
	"net/http"
)

//Main entry point for the API Gateway
func main() {
	//Register handler for root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gateway!\n")
		fmt.Fprintf(w, "You requested: %s\n", r.URL.Path)
		fmt.Fprintf(w, "Method: %s\n", r.Method)

	})

	//Register handler for health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status": "healthy"}`)
	})

	//Start the server
	addr := "0.0.0.0:8080"
	log.Printf("Server starting on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}