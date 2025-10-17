MIGRATION_PATH = ./cmd/migrations
BINARY_NAME = main
include app.env

## Database Migrations

.PHONY: migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATION_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDRESS) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDRESS) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-status
migrate-status:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDRESS) version

.PHONY: migrate-force
migrate-force:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDRESS) force $(filter-out $@,$(MAKECMDGOALS))

## Application Build & Run

.PHONY: build
build:
	@echo "Building Backend..."
	@go build -o ${BINARY_NAME} .
	@echo "Binary built: ${BINARY_NAME}"

.PHONY: run
run: build
	@echo "Starting Backend..."
	@DSN="${DB_ADDRESS}" ENV="${ENVIRONMENT}" ./${BINARY_NAME}

.PHONY: dev
dev:
	@echo "Starting Backend in development mode..."
	@DSN="${DB_ADDRESS}" ENV="${ENVIRONMENT}" go run .

.PHONY: clean
clean:
	@echo "Cleaning..."
	@go clean
	@rm -f ${BINARY_NAME}
	@echo "Cleaned"

.PHONY: start
start: run

.PHONY: stop
stop:
	@echo "Stopping Backend..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}" || true
	@echo "Stopped Backend"

.PHONY: restart
restart: stop run

## Docker Commands

.PHONY: docker-up
docker-up:
	@echo "Starting Docker containers..."
	docker compose up -d

.PHONY: docker-down
docker-down:
	@echo "Stopping Docker containers..."
	docker compose down

.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker compose build

.PHONY: docker-up-log
docker-up-log:
	@echo "Docker Up..."
	docker compose up

.PHONY: docker-logs
docker-logs:
	docker compose logs -f app

.PHONY: docker-restart
docker-restart: docker-down docker-up

## Development Utilities

.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod verify

.PHONY: tidy
tidy:
	@echo "Tidy dependencies..."
	go mod tidy

.PHONY: test
test:
	@echo "Running tests..."
	go test ./... -v

.PHONY: watch
watch:
	@echo "Watching for changes..."
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "Installing air..." && \
		go install github.com/cosmtrek/air@latest && \
		air; \
	fi

## Help
.PHONY: help
help:
	@echo "Available commands:"
	@echo ""
	@echo "Database Migrations:"
	@echo "  migrate-create <name>  Create new migration files"
	@echo "  migrate-up            Apply all pending migrations"
	@echo "  migrate-down [steps]  Rollback migrations (optional steps)"
	@echo "  migrate-status        Show current migration version"
	@echo "  migrate-force <version> Force set migration version"
	@echo ""
	@echo "Application:"
	@echo "  build                 Build the application"
	@echo "  run                   Build and run the application"
	@echo "  dev                   Run with hot-reload using air"
	@echo "  start                 Alias for run"
	@echo "  stop                  Stop the running application"
	@echo "  restart               Restart the application"
	@echo "  clean                 Remove built binaries"
	@echo ""
	@echo "Docker:"
	@echo "  docker-up            Start Docker containers"
	@echo "  docker-down          Stop Docker containers"
	@echo "  docker-build         Build Docker image"
	@echo "  docker-logs          Show application logs"
	@echo "  docker-restart       Restart Docker containers"
	@echo ""
	@echo "Development:"
	@echo "  deps                 Download dependencies"
	@echo "  tidy                 Tidy go.mod"
	@echo "  test                 Run tests"
	@echo "  watch                Run with hot-reload"
	@echo ""
	@echo "  help                 Show this help message"

.DEFAULT_GOAL := help