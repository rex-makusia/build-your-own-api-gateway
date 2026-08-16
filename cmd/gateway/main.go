package main

import (
    "log"
    "net/http"
    "build-your-own-api-gateway/internal/gateway"
    "build-your-own-api-gateway/api"
    "build-your-own-api-gateway/internal/middleware"
)

func main() {
    // Initialize the server
    srv := gateway.NewServer()

    // Set up middleware
    srv.Use(middleware.AuthMiddleware)

    // Set up routes
    api.RegisterRoutes(srv)

    // Start the server
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", srv); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}