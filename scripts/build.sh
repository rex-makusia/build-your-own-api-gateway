#!/bin/bash

# Build the Go application
go build -o bin/gateway ./cmd/gateway

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo "Build successful! The binary is located at ./bin/gateway"
else
    echo "Build failed. Please check the errors above."
fi