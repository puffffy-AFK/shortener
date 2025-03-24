FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o shortener ./cmd/http_server

FROM alpine:latest

WORKDIR /app

RUN mkdir -p /app

COPY --from=builder /app/shortener /app/
RUN chmod +x /app/shortener

COPY config.yaml /app/

ENTRYPOINT ["/app/shortener"]