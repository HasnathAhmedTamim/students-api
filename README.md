# Student API

A RESTful API built in **Go** for managing student records using **SQLite** as the database.

## Tech Stack

| Component | Technology |
|-----------|------------|
| Language | Go |
| Database | SQLite (pure Go driver: `modernc.org/sqlite`) |
| HTTP Router | Standard `net/http` (Go 1.22+) |
| Validation | `go-playground/validator` |
| Config | YAML configuration |

## Project Structure

```
student-api/
├── cmd/students-api/
│   └── main.go              # Application entry point
├── internal/
│   ├── config/              # Configuration loading
│   ├── http/handlers/student/ # HTTP handlers
│   ├── storage/             # Storage interface
│   │   └── sqlite/          # SQLite implementation
│   ├── types/               # Data types (Student struct)
│   └── utils/response/      # JSON response helpers
└── config/
    └── local.yaml           # Configuration file
```

## Getting Started

### Prerequisites

- Go 1.22 or higher

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/hasnathahmedtamim/students-api.git
   cd student-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run cmd/students-api/main.go -config config/local.yaml
   ```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/students` | Create a new student |
| `GET` | `/api/students/{id}` | Get student by ID |
| `GET` | `/api/students` | Get all students |
| `PUT` | `/api/students/{id}` | Update student by ID |
| `DELETE` | `/api/students/{id}` | Delete student by ID |

## Student Model

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "age": 25
}
```

### Validation Rules

| Field | Rules |
|-------|-------|
| `name` | Required |
| `email` | Required, valid email format |
| `age` | Required |

## API Usage Examples

### Create Student

```bash
curl -X POST http://localhost:8080/api/students \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","age":25}'
```

**Response:**
```json
{
  "success": "OK",
  "id": 1
}
```

### Get All Students

```bash
curl http://localhost:8080/api/students
```

**Response:**
```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "age": 25
  }
]
```

### Get Student by ID

```bash
curl http://localhost:8080/api/students/1
```

**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "age": 25
}
```

### Update Student

```bash
curl -X PUT http://localhost:8080/api/students/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"John Updated","email":"john@example.com","age":26}'
```

**Response:**
```json
{
  "success": "OK"
}
```

### Delete Student

```bash
curl -X DELETE http://localhost:8080/api/students/1
```

**Response:**
```json
{
  "success": "OK"
}
```

## Features

- **CRUD Operations** - Full create, read, update, delete functionality
- **Request Validation** - Using `go-playground/validator`
- **Graceful Shutdown** - Server handles `SIGINT`/`SIGTERM` signals
- **Structured Logging** - Using Go's `slog` package
- **Configuration Management** - YAML-based configuration
- **Clean Architecture** - Separation of concerns with interfaces

## Configuration

Create a `config/local.yaml` file:

```yaml
env: "local"
storage_path: "./storage.db"
http_server:
  address: "localhost:8080"
```

## License

MIT License