# URL Shortener

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-316192?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

---

## About This Project

This is an implementation of the **URL Shortener**, inspired by [ByteByteGo's System Design Interview: Design A URL Shortener](https://bytebytego.com/courses/system-design-interview/design-a-url-shortener). The goal is to understand system design concepts by building core components from scratch.

---

## Tech Stack

- **Language**: Go 1.24+
- **Web Framework**: Gin
- **Database**: PostgreSQL 15
- **Containerization**: Docker & Docker Compose

---

## Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/pongpradk/url-shortener.git
cd url-shortener
```

### 2. Configure Environment Variables

```bash
cp .env.example .env
```

Update .env if needed (default values work for local development).

### 3. Start PostgreSQL

```bash
docker compose up -d

```

### 4. Install Go Dependencies

```bash
go mod download
```

### 5. Run the Application

```bash
go run cmd/server/main.go
```

---

## API Documentation

### 1. Shorten URL

Create a short URL from a long URL.

**Request**

```http
POST /api/v1/data/shorten
Content-Type: application/json

{
  "longUrl": "https://github.com/pongpradk/url-shortener"
}
```

**Response**

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "shortUrl": "79ng5VsdJ6s"
}
```

**cURL Example**

```bash
curl -X POST http://localhost:8080/api/v1/data/shorten \
  -H "Content-Type: application/json" \
  -d '{"longUrl": "https://github.com/pongpradk/url-shortener"}'
```

### 2. Redirect to Original URL

**Request**

```http
GET /{shortUrl}
```

**Response**

```http
Location: https://github.com/pongpradk/url-shortener
Status: 301 Moved Permanently
```

**Browser/cURL Example**

```bash
curl -L http://localhost:8080/79ng5VsdJ6s
```

---

## Project Structure

```
url-shortener-go/
├── cmd/                         # Command-related files
│   └── server/                  # Application entry point
│       └── main.go              # Main application logic
├── internal/                    # Internal codebase
│   ├── database/                # Database connection setup
│   │   └── database.go
│   ├── encoder/                 # Base62 encoding
│   │   └── base62.go
│   ├── handler/                 # HTTP request handlers (controllers)
│   │   └── url_handler.go
│   ├── repository/              # Data access (repository layer)
│   │   ├── url_repository.go
│   └── service/                 # Business logic (service layer)
│       └── url_service.go
├── .env.example
├── .gitignore
├── docker-compose.yml
├── schema.sql                   # Database schema
├── go.mod                       # Go module dependencies
├── go.sum                       # Go module checksum file
└── README.md
```

---

## License

This project is licensed under the MIT License.

---
