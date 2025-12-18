# SSL Checker Backend (Go)

This is the core logic of the application, written in Go. It follows a standard Go project layout to ensure maintainability and separation of concerns.

## 📂 Project Structure

```text
backend/
├── cmd/
│   ├── api/            # Entry point for the HTTP Server (main.go)
│   └── cli/            # Entry point for the CLI tool (main.go)
├── internal/
│   ├── handlers/       # HTTP Controllers (API endpoints)
│   └── ssllabs/        # Core Business Logic (SSL Labs Client)
├── go.mod              # Module definition
└── Dockerfile          # Container definition
```

## ⚙️ How to Run Locally (Without Docker)

If you prefer to run Go directly on your machine:

### 1. Run the HTTP Server
This starts the REST API used by the frontend.
```bash
go run cmd/api/main.go
# Server listens on http://localhost:8080
```

### 2. Run the CLI Tool
A standalone command-line interface to check domains directly from the terminal.
```bash
# Basic usage (uses cache if available)
go run cmd/cli/main.go google.com

# Force a fresh scan (async response)
go run cmd/cli/main.go google.com new

# Force a fresh scan (sync response with pooling)
go run cmd/cli/main.go google.com newsync
```