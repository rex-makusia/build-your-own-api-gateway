package gateway

import (
    "net/http"
)

// HelloHandler responds with a simple greeting.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello, World!"))
}

// HealthCheckHandler checks the health of the API gateway.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("API Gateway is up and running"))
}