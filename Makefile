.PHONY: help build run test clean deps docker-up docker-down

# Default target
help:
	@echo "Available commands:"
	@echo "  build      - Build the application"
	@echo "  run        - Run the application"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean build artifacts"
	@echo "  deps       - Install dependencies"
	@echo "  docker-up  - Start PostgreSQL with Docker Compose"
	@echo "  docker-down- Stop Docker Compose services"
	@echo "  setup      - Setup development environment"

# Build the application
build:
	@echo "Building application..."
	go build -o bin/gitlab-tool main.go

# Run the application
run:
	@echo "Running application..."
	go run main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	go clean

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

# Start PostgreSQL with Docker Compose
docker-up:
	@echo "Starting PostgreSQL..."
	docker-compose up -d postgres
	@echo "PostgreSQL is running on localhost:5432"
	@echo "pgAdmin is available at http://localhost:8081 (admin@example.com / admin)"

# Stop Docker Compose services
docker-down:
	@echo "Stopping services..."
	docker-compose down

# Setup development environment
setup: docker-up deps
	@echo "Development environment setup complete!"
	@echo "PostgreSQL is running on localhost:5432"
	@echo "Run 'make run' to start the application"

# Development mode with hot reload (requires air: go install github.com/cosmtrek/air@latest)
dev:
	@echo "Starting development mode with hot reload..."
	air

# Database migrations
migrate:
	@echo "Running database migrations..."
	go run main.go migrate

# Generate API documentation (requires swag: go install github.com/swaggo/swag/cmd/swag@latest)
docs:
	@echo "Generating API documentation..."
	swag init -g main.go

# Lint code
lint:
	@echo "Linting code..."
	golangci-lint run

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...
	go vet ./... 