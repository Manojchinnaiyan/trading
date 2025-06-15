# Makefile
.PHONY: help build run test clean docker-build docker-up docker-down proto

# Default target
help:
	@echo "Available commands:"
	@echo "  build         Build the Go application"
	@echo "  run           Run the application locally"
	@echo "  test          Run unit tests"
	@echo "  clean         Clean build artifacts"
	@echo "  docker-build  Build Docker image"
	@echo "  docker-up     Start all services with Docker Compose"
	@echo "  docker-down   Stop all services"
	@echo "  proto         Generate gRPC code from proto files"
	@echo "  setup         Initial setup (install dependencies)"

# Build the Go application
build:
	@echo "Building application..."
	go build -o bin/trading-platform-backend .

# Run the application locally
run:
	@echo "Running application..."
	go run main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	go clean

# Setup development environment
setup:
	@echo "Setting up development environment..."
	go mod download
	go mod tidy

# Docker commands
docker-build:
	@echo "Building Docker image..."
	docker build -t trading-platform-backend:latest .

docker-up:
	@echo "Starting services with Docker Compose..."
	docker-compose up -d

docker-down:
	@echo "Stopping services..."
	docker-compose down

docker-logs:
	@echo "Showing logs..."
	docker-compose logs -f

# Generate gRPC code (bonus feature)
proto:
	@echo "Generating gRPC code..."
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
		proto/trading.proto

# Development helpers
dev-db:
	@echo "Starting only database services..."
	docker-compose up -d postgres redis

dev-stop:
	@echo "Stopping development services..."
	docker-compose stop

# Production deployment
deploy:
	@echo "Deploying to production..."
	docker-compose -f docker-compose.yml up -d

# Database migrations (if needed)
migrate-up:
	@echo "Running database migrations..."
	# Add migration commands here

migrate-down:
	@echo "Rolling back database migrations..."
	# Add rollback commands here