.PHONY: run build fmt lint test migrate deps docker-up docker-down db-create

# Переменные
BIN=shortener
GO=go
GOMOD=$(GO) mod
DB_HOST=localhost
DB_USER=user
DB_PASSWORD=password
DB_NAME=shortener

# Установка зависимостей
deps:
	$(GOMOD) tidy

# Форматирование кода
fmt:
	$(GO) fmt ./...

# Линтер (golangci-lint)
lint:
	golangci-lint run

# Сборка бинарника
build:
	$(GO) build -o $(BIN) cmd/http_server/main.go

# Создание базы данных (если ее нет)
db-create:
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -U $(DB_USER) -d postgres \
	-tc "SELECT 1 FROM pg_database WHERE datname='$(DB_NAME)'" | grep -q 1 || \
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -U $(DB_USER) -d postgres \
	-c "CREATE DATABASE $(DB_NAME);"
# Миграции
migrate:
	psql postgres://user:password@localhost/shortener?sslmode=disable -f internal/client/db/migrations/001_init.sql

# Запуск сервера (автоматически собирает бинарник, создает БД и запускает миграции)
run: build db-create migrate
	./$(BIN)

# Запуск контейнеров
docker-up:
	docker-compose up -d

# Остановка контейнеров
docker-down:
	docker-compose down