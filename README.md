# URL Shortener

[![Go Version](https://img.shields.io/badge/Go-1.24.11-00ADD8?style=flat&logo=go)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-316192?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A URL shortener service implementing concepts from **"System Design Interview"** by Alex Xu - ByteByteGo

---

## About This Project

This is an implementation of the **URL Shortener** system adapted from [ByteByteGo's System Design Interview: Design A URL Shortener](https://bytebytego.com/courses/system-design-interview/design-a-url-shortener). The goal is to deeply understand the system design concept by building them from scratch.

---

## Architecture

```
        ┌─────────────┐
        │   Client    │
        └──────┬──────┘
               │ HTTP Request
               ↓
┌─────────────────────────────────┐
│       Gin HTTP Handler          │
│  (URL validation & routing)     │
└──────────────┬──────────────────┘
               │
               ↓
┌─────────────────────────────────┐
│      Business Logic Layer       │
│  (URL shortening & retrieval)   │
└──────────────┬──────────────────┘
               │
               ↓
┌─────────────────────────────────┐
│      Repository Layer           │
│  (Database operations)          │
└──────────────┬──────────────────┘
               │
               ↓
┌─────────────────────────────────┐
│      PostgreSQL Database        │
│      (Persistent storage)       │
└─────────────────────────────────┘
```

---

## Tech Stack

- **Language**: [Go 1.24.11](https://go.dev/)
- **Web Framework**: [Gin](https://gin-gonic.com/)
- **Database**: [PostgreSQL 15](https://www.postgresql.org/)
- **Containerization**: [Docker](https://www.docker.com/)

---

## Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/pongpradk/url-shortener.git
cd url-shortener
```

### 2. Configure Environment Variables

Copy the example environment file

```bash
cp .env.example .env

# Edit .env if needed (defaults work for local development)
```

### 3. Start PostgreSQL Database

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

### Endpoints

#### 1. Shorten URL

Create a short URL from a long URL.

**Request:**

```http
POST /api/v1/data/shorten
Content-Type: application/json

{
  "longUrl": "https://github.com/pongpradk/url-shortener"
}
```

**Response:**

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "shortUrl": "79ng5VsdJ6s"
}
```

**cURL Example:**

```bash
curl -X POST http://localhost:8080/api/v1/data/shorten \
  -H "Content-Type: application/json" \
  -d '{"longUrl": "https://github.com/pongpradk/url-shortener"}'
```

#### 2. Redirect to Original URL

Access the short URL to be redirected to the original URL.

**Request:**

```http
GET /{shortUrl}
```

**Response:**

```http
HTTP/1.1 301 Moved Permanently
Location: https://github.com/pongpradk/url-shortener
```

**Browser/cURL Example:**

```bash
curl -L http://localhost:8080/79ng5VsdJ6s
# Redirects to the original URL
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
│   ├── repository/              # Database operations (repository layer)
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
