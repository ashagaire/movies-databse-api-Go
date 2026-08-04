# movies-api

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