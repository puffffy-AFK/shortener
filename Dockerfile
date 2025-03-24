# Stage 1: Build the Go app
FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Скомпилируем файл shortener для Linux
RUN GOOS=linux GOARCH=amd64 go build -o shortener ./cmd/http_server

# Stage 2: Run the Go app
FROM alpine:latest

WORKDIR /app

# Создаем директорию /app
RUN mkdir -p /app

# Копируем собранный файл из builder стадии
COPY --from=builder /app/shortener /app/
RUN chmod +x /app/shortener

# Копируем файл config.yaml
COPY config.yaml /app/

ENTRYPOINT ["/app/shortener"]