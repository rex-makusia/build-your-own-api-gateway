# Makefile for build-your-own-api-gateway project

.PHONY: build run clean

build:
	go build -o bin/gateway ./cmd/gateway

run: build
	./bin/gateway

clean:
	go clean
	rm -rf bin/gateway