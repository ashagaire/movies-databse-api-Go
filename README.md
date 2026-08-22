# movies-api

A high performance, relational REST API built in Go for managing structured movie metadata. Designed with a strict layered architecture, this API provides a strong foundation for applications requiring complex data relationships, dynamic querying, and strict data integrity.

## Architecture & Design

This project implements a clean, layered architecture to ensure separation of concerns, maintainability, and testability:

*   **Handler Layer:** Manages HTTP request parsing, routing, and JSON serialization.
*   **Service Layer:** Encapsulates business logic, data validation, and transactional orchestration.
*   **Repository Layer:** Interfaces directly with the SQLite database, executing raw SQL queries and managing complex `JOIN` operations.

## Core Capabilities

*   **Relational Data Management:** Full CRUD support for Movies, Actors, and Genres.
*   **Complex Associations:** Handles Many-to-Many relationships (Movies ↔ Genres, Movies ↔ Actors) via junction tables.
*   **Dynamic Querying:** 
    *   Pagination (`?page=0&size=10`)
    *   Attribute filtering (`?year=1999&genre=1`)
    *   Case-insensitive partial text search (`/search?title=matrix`)
*   **Strict Data Integrity:** Enforces `ON DELETE RESTRICT` foreign key constraints at the database level to prevent orphaned records.
*   **Transactional Operations:** Implements database transactions (`BEGIN`, `COMMIT`, `ROLLBACK`) for multi-table inserts and safe `?force=true` cascading deletions.

## Tech Stack

*   **Language:** Go (1.22+)
*   **Database:** SQLite3 (`github.com/mattn/go-sqlite3`)
*   **Routing:** Go Standard Library (`net/http` )

## Setup and Installation


### Local Development

Ensure Go 1.22+ and a C compiler (required for `go-sqlite3` CGO bindings) are installed.

1. Clone the repository:
   ```bash
   git clone https://gitea.kood.tech/muhammadowaisjaved/movies-api.git
   cd movies-api
    ```
2. Download dependencies:
    ```bash
    go mod tidy
    ```

3. Run the application:
    
    ```bash
    go run ./cmd/api/main.go
    ```

## Containerization

The application includes a multi-stage Dockerfile and docker-compose.yml for consistent deployment across environments.
    
    ```
    docker-compose up --build
    ```

The API will be exposed on http://localhost:8080. The SQLite database file is persisted via a Docker volume in the ./data directory.

## Project Structure

```text

open-movies-db/
├── cmd/
│   └── api/
│       └── main.go           # Entry point: wires up dependencies and starts the server
├── internal/
│   ├── models/               # Struct definitions (Movie, Actor, Genre)
│   ├── database/             # SQLite connection setup (db.go)
│   ├── repository/           # Raw SQL queries (genre.go, movie.go, actor.go)
│   ├── service/              # Logic & validation (genre.go, movie.go, actor.go)
│   └── handler/              # HTTP routing, JSON encoding/decoding (genre.go, movie.go, actor.go)
├── pkg/
│   └── errors/               # Custom error definitions (e.g., ErrNotFound)
├── data/                     # Local folder for the SQLite .db file (ignored in git)
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum

```
