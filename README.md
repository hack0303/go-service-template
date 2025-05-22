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
├── test/             # Test files
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

4. Run the application:
   ```bash
   go run main.go
   ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
