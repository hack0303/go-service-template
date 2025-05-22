# Makefile for Go Service Template

.PHONY: build run test clean lint

# Build the application
build:
	go build -o bin/main cmd/main.go

# Run the application
run:
	go run cmd/main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Lint the code
lint:
	golangci-lint run

# Generate Swagger docs
swagger:
	swag init -g cmd/main.go

# Run with hot reload using air
dev:
	air

# Run database migrations
migrate:
	go run cmd/migrate.go

# Build Docker image
docker-build:
	docker build -t go-service-template .

# Run Docker container
docker-run:
	docker run -p 8080:8080 go-service-template
