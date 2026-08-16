FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o gateway ./cmd/gateway/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/gateway .

CMD ["./gateway"]