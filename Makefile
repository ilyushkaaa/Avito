include .env
export


MIGRATION_FOLDER=$(CURDIR)/internal/database/postgres/migrations

.PHONY: app_start
app_start:
	docker-compose up --build

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) host=localhost port=$(DB_PORT) sslmode=disable" up

.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) host=localhost port=$(DB_PORT) sslmode=disable" down

.PHONY: build
build:
	go build cmd/main/main.go