include .env
export 

.PHONY: run test tidy lint


run:
	air

tui:
	go run ./cmd/tui

test:
	go test ./... -count=1

tidy:
	go mod tidy

build:
	go build -o bin/api ./cmd/api

start:
	docker-compose up -d
stop:
	docker-compose down -v

#init db and applies migrations
db-migrate:
	migrate -path migrations -database "$$DB_URL" up

# Rollback the last migration
db-migrate-down:
	migrate -path migrations -database "$$DB_URL" down 1

db-reset:
	migrate -path migrations -database "$$DB_URL" drop -f
	migrate -path migrations -database "$$DB_URL" up