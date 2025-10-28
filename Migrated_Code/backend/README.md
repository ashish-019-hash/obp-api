# OBP-API Backend Service

A Golang backend service based on the Open Bank Project API analysis from phase-01-output.

## Project Structure

```
/backend
├── cmd/                    # Application entry point
│   └── main.go
├── internal/               # Private application code
│   ├── config/            # Configuration loading
│   ├── controllers/       # HTTP request handlers
│   ├── middleware/        # Custom middleware
│   ├── models/           # Entity definitions
│   ├── repositories/     # Database access layer
│   ├── routes/           # API route registration
│   ├── services/         # Business logic layer
│   └── utils/            # Utility functions
├── pkg/                   # Reusable packages
│   └── db/               # Database connection helpers
├── test/                  # Test files
├── go.mod                # Go module definition
└── go.sum                # Go module checksums
```

## Getting Started

### Prerequisites
- Go 1.21 or higher
- Git

### Installation

1. Clone the repository
2. Navigate to the backend directory
3. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running the Server

```bash
go run cmd/main.go
```

The server will start on port 8080 (configurable via PORT environment variable).

### Available Endpoints

- `GET /health` - Health check endpoint
- `GET /ping` - Simple ping endpoint
- `GET /api/v1/health` - API health check

### Testing

Run tests with:
```bash
go test ./test/...
```

### Configuration

Configuration is loaded from environment variables. See `.env` file for available options.

## Based on OBP-API Analysis

This backend service is built based on comprehensive analysis of the Open Bank Project API:

- **Screen Flows**: 12 documented user journeys and navigation patterns
- **Entities**: 113+ business entities across banking domains
- **Business Rules**: 12 core calculation and decision rules
- **Validation Rules**: 38 input validation and data integrity rules
- **User Stories**: 50+ business-level user stories across 10 domains

## Development Status

**Step 1: Project Skeleton** ✅ Complete
- Go project structure initialized
- Basic HTTP server with Gin framework
- Configuration management
- Health check endpoints

**Next Steps**: Awaiting user input for Step 2 implementation.
