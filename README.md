# movies-api

A high performance, relational REST API built in Go for managing structured movie metadata. Designed with a strict layered architecture, this API provides a strong foundation for applications requiring complex data relationships, dynamic querying, and strict data integrity.

## Architecture & Design

This project implements a clean, layered architecture to ensure separation of concerns, maintainability, and testability:

*   **Handler Layer:** Manages HTTP request parsing, routing, and JSON serialization.
*   **Service Layer:** Encapsulates business logic, data validation, and transactional orchestration.
*   **Repository Layer:** Interfaces directly with the SQLite database, executing raw SQL queries and managing complex `JOIN` operations.

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
