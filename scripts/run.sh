#!/bin/bash

# Navigate to the project directory
cd "$(dirname "$0")/.."

# Run the application
go run ./cmd/gateway/main.go