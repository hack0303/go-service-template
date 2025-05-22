# Go Service Template

A production-ready Go service template implementing Clean Architecture with best practices.

## Clean Architecture Overview

This project follows Clean Architecture principles, separating the system into layers with clear boundaries and dependencies:

1. **Entities**: Core business objects and rules
2. **Use Cases**: Application-specific business rules
3. **Interface Adapters**: Controllers, presenters, and gateways
4. **Frameworks & Drivers**: Web, UI, DB, and external interfaces

Dependencies flow inward, with outer layers depending on inner layers.

## Project Structure

```
.
├── cmd/              # Main application and command line interface
├── internal/         # Private application and library code
│   ├── entity/       # Business entities and rules
│   ├── usecase/      # Application business logic
│   ├── handler/      # Interface adapters (HTTP handlers)
│   ├── repository/   # Interface adapters (data access)
│   ├── middleware/   # Interface adapters (HTTP middleware)
│   └── config/       # Configuration management
├── pkg/              # Public library code
├── migrations/       # Database migrations
├── scripts/          # Build and deployment scripts
├── doc/              # Documentation
│   ├── config/       # Configuration files for deployments
│   ├── sql/          # SQL scripts for deployments
│   └── feat/         # Feature documentation
├── test/             # Test files
│   ├── http/         # REST client test scripts (*.http)
├── go.mod            # Go module dependencies
├── main.go           # Application entry point
└── README.md         # Project documentation
```

## Technology Stack

- **Language**: Go 1.21+
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Configuration**: Viper
- **Logging**: Zap
- **Testing**: Testify
- **API Documentation**: Swagger
- **Containerization**: Docker
- **CI/CD**: GitHub Actions

## HTTP Message Structure Standards

### Message Body Structure

HTTP response messages use JSON format with the following fields:

- **code**: Integer, status code indicating request processing result
- **msg**: String, describing request processing status or error message
- **data**: JSON object, carrying response data (optional)

### Field Definitions

| Field | Type | Description | Required |
|-------|------|-------------|----------|
| code  | Integer | Status code indicating result | Yes |
| msg   | String | Status description or error message | Yes |
| data  | Object | Response data, structure defined by business needs | No |

### Status Code Design

Status codes are 11-digit integers following the structure: `{exception category:3}{exception system:4}{exception custom:4}`

#### Status Code Rules

- **Success**:
  - `200`: Request successful
- **Failure**:
  - `400` or `400XXXXYYYY`: Parameter-related errors
  - `500` or `500XXXXYYYY`: System-related errors

#### Status Code Validation

- **Range Check**: Ensures status code is within integer range
- **Structure Check**:
  - Exception category must be 200, 400, or 500
  - Exception system and custom parts must be unique

### Examples

#### Success Response
```json
{
  "code": 200,
  "msg": "Success",
  "data": {
    "user_id": 123,
    "username": "example"
  }
}
```

#### Parameter Error
```json
{
  "code": 40010010001,
  "msg": "Invalid parameter: user_id is missing",
  "data": {
    "error_field": "user_id",
    "error_detail": "required field"
  }
}
```

#### System Error
```json
{
  "code": 50020030002,
  "msg": "Database connection failed",
  "data": {}
}
```

### Implementation Details

The HTTP response structure is implemented in `internal/handler/response.go` with:
- `HTTPResponse` struct
- `NewSuccessResponse` and `NewErrorResponse` constructors
- `WriteResponse` helper function
- `ValidateStatusCode` validation function

## Components

### Core Components
- **Configuration Management**: Centralized configuration with environment variable support
- **Logging**: Structured logging with Zap
- **Database**: PostgreSQL with GORM ORM
- **HTTP Server**: Gin web framework with middleware support
- **Health Checks**: Built-in health check endpoints
- **Metrics**: Prometheus metrics endpoint

### Optional Components
- **Authentication**: JWT-based authentication
- **Caching**: Redis integration
- **Message Queue**: RabbitMQ support
- **Email Service**: SMTP email integration

## Docker Setup

### Development with Docker Compose

1. Start the development environment:
   ```bash
   docker-compose up
   ```

2. Access the application:
   ```bash
   http://localhost:8080
   ```

3. Access the database:
   ```bash
   psql postgres://user:password@localhost:5432/go_service
   ```

### Production Build

1. Build the Docker image:
   ```bash
   docker build -t go-service-template .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 go-service-template
   ```

## Makefile Commands

The project includes a Makefile with common development tasks:

```bash
# Build the application
make build

# Run the application
make run

# Run tests
make test

# Clean build artifacts
make clean

# Lint the code
make lint

# Generate Swagger docs
make swagger

# Run with hot reload using air
make dev

# Run database migrations
make migrate

# Build Docker image
make docker-build

# Run Docker container
make docker-run
```

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/go-service-template.git
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Configure environment variables:
   ```bash
   cp .env.example .env
   ```

4. Run the application (choose one):
   ```bash
   # Using make
   make run

   # Directly
   go run cmd/main.go
   ```

5. For development with hot reload:
   ```bash
   make dev
   ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
