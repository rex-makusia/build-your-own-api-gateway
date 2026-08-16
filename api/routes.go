package api

import (
    "net/http"
    "github.com/gorilla/mux"
    "your_project/internal/gateway" // Adjust the import path according to your project structure
)

func RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/api/your-endpoint", gateway.YourHandlerFunction).Methods(http.MethodGet)
    // Add more routes here
}