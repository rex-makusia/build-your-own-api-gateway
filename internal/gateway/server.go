package main

import (
    "log"
    "net/http"
)

func StartServer(addr string) {
    mux := http.NewServeMux()

    // Here you would register your routes and middleware
    // For example: mux.HandleFunc("/example", exampleHandler)

    log.Printf("Starting server on %s", addr)
    if err := http.ListenAndServe(addr, mux); err != nil {
        log.Fatalf("Could not start server: %s", err)
    }
}