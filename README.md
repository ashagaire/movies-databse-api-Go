# movies-api

![Banner](./static/assets/banner.png)

A high performance, relational REST API built in Go for managing structured movie metadata. Designed with a strict layered architecture, this API provides a strong foundation for applications requiring complex data relationships, dynamic querying, and strict data integrity.

## Architecture & Design

This project implements a clean, layered architecture to ensure separation of concerns, maintainability, and testability:

- **Handler Layer:** Manages HTTP request parsing, routing, and JSON serialization.
- **Service Layer:** Encapsulates business logic, data validation, and transactional orchestration.
- **Repository Layer:** Interfaces directly with the SQLite database, executing raw SQL queries and managing complex `JOIN` operations.

## Core Capabilities

- **Relational Data Management:** Full CRUD support for Movies, Actors, and Genres.
- **Complex Associations:** Handles Many-to-Many relationships (Movies ↔ Genres, Movies ↔ Actors) via junction tables.
- **Dynamic Querying:**
  - Case-insensitive partial text search (`/search?title=matrix`)
- **Strict Data Integrity:** Enforces `ON DELETE RESTRICT` foreign key constraints at the database level to prevent orphaned records.
- **Transactional Operations:** Implements database transactions (`BEGIN`, `COMMIT`, `ROLLBACK`) for multi-table inserts and safe `?force=true` cascading deletions.

## Tech Stack

- **Language:** Go (1.22+)
- **Database:** SQLite3 (`github.com/mattn/go-sqlite3`)
- **Routing:** Go Standard Library (`net/http` )
- **Migrations:** Goose
- **Containerization:** Docker & Docker Compose
- **Testing:** Go testing with isolated in-memory SQLite databases
- **API Testing:** Postman or Curl

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

3. Database Setup and migration:

   install the Goose CLI from the system terminal instead of VSCode/Zed terminal:

   ```bash
   go install github.com/pressly/goose/v3/cmd/goose@latest
   ```

    Ensure your Go binary path is included in your shell environment (PATH):

    ```bash
    # For Zsh (macOS default / Linux)
    echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc && source ~/.zshrc


    # For Bash
    echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc && source ~/.bashrc
   ```

   Verify the installation:

   ```bash
   goose -version
   ```

   Create the database and run migrations:

   ```bash
   goose -dir migrations sqlite3 ./internal/database/movies.db up
   ```

4. Run the application:

   ```bash
   go run ./cmd/api/main.go
   ```

## Containerization

The application includes a multi-stage Dockerfile and docker-compose.yml for consistent deployment across environments.

    docker-compose up --build

The API will be exposed on http://localhost:8080. The SQLite database file is persisted via a Docker volume in the ./data directory.

## Testing

The project includes a comprehensive automated testing suite. Tests are executed against isolated, in-memory SQLite databases (:memory:) to ensure no impact on the local development environment.

To run the unit tests in the root directory:

    go test ./... -v

The project also includes Postman collection for basic CRUD APIs. You can import the collection file to Postman application and test apis and their response status.

## API Reference

### Movies

| Method   | Endpoint             | Description                                                                  |
| :------- | :------------------- | :--------------------------------------------------------------------------- |
| `POST`   | `/api/movies`        | Create a movie and establish actor/genre relationships                       |
| `GET`    | `/api/movies`        | Retrieve movies (Supports `size`, `year`, `genre` and `actor` query params)  |
| `GET`    | `/api/movies/search` | Search movies by title (`?title=`)                                           |
| `GET`    | `/api/movies/{id}`   | Retrieve a specific movie                                                    |
| `PATCH`  | `/api/movies/{id}`   | Partially update movie attributes or relationships                           |
| `DELETE` | `/api/movies/{id}`   | Delete a movie (Fails if relationships exist unless `?force=true` is passed) |

### Genres & Actors

| Method   | Endpoint                                 | Description                               |
| :------- | :--------------------------------------- | :---------------------------------------- |
| `POST`   | `/api/genres` or `/api/actors`           | Create a new entity                       |
| `GET`    | `/api/genres` or `/api/actors`           | Retrieve all entities                     |
| `GET`    | `/api/genres/{id}` or `/api/actors/{id}` | Retrieve a specific entity                |
| `PATCH`  | `/api/genres/{id}` or `/api/actors/{id}` | Partially update an entity                |
| `DELETE` | `/api/genres/{id}` or `/api/actors/{id}` | Delete an entity (Supports `?force=true`) |


## Usage & Payload Examples

Below are examples of how to interact with the API using curl. You can use these JSON payloads in Postman as well.

### 1. Genres

Create a Genre:

```Bash
curl -X POST http://localhost:8080/api/genres \
-H "Content-Type: application/json" \
-d '{
  "name": "Action"
}'
```

Update a Genre (PATCH ):

```Bash
curl -X PATCH http://localhost:8080/api/genres/1 \
-H "Content-Type: application/json" \
-d '{
  "name": "Action & Adventure"
}'
```

### 2. Actors

Create an Actor:

_(Note: birthDate must strictly follow the ISO 8601 YYYY-MM-DD format )_

```Bash
curl -X POST http://localhost:8080/api/actors \
-H "Content-Type: application/json" \
-d '{
  "name": "Leonardo DiCaprio",
  "birthDate": "1974-11-11"
}'
```

Update an Actor (PATCH ):

```Bash
curl -X PATCH http://localhost:8080/api/actors/1 \
-H "Content-Type: application/json" \
-d '{
  "name": "Leo DiCaprio"
}'
```

### 3. Movies

Create a Movie with Relationships:

_(Note: The genres and actors arrays expect objects containing the id of existing records )_

```Bash
curl -X POST http://localhost:8080/api/movies \
-H "Content-Type: application/json" \
-d '{
  "title": "Inception",
  "releaseYear": 2010,
  "duration": 148,
  "genres": [{"id": 1}, {"id": 2}],
  "actors": [{"id": 1}, {"id": 3}]
}'
```

Update a Movie (PATCH ):

_(You can send just one field, or multiple. Providing a new array for relationships will overwrite the old ones)_

```Bash
curl -X PATCH http://localhost:8080/api/movies/1 \
-H "Content-Type: application/json" \
-d '{
  "duration": 150,
  "actors": [{"id": 1}]
}'
```

### 4. Filtering & Searching (GET )

The API supports dynamic querying via URL parameters.

Search by Partial Name/Title:

```Bash
curl -X GET "http://localhost:8080/api/movies/search?title=incept"
curl -X GET "http://localhost:8080/api/actors?name=Leo"
```

Filter Movies by Relationships or Year:

```Bash
curl -X GET "http://localhost:8080/api/movies?genre=1"
curl -X GET "http://localhost:8080/api/movies?actor=3"
curl -X GET "http://localhost:8080/api/movies?year=2010"
```

### 5. Deletions & Constraints

By default, deleting an entity with existing relationships will return a 400 Bad Request to protect data integrity.

Standard Delete (Will fail if relationships exist ):

```Bash
curl -X DELETE http://localhost:8080/api/genres/1
```

Force Delete (Cascades and removes relationships ):

```Bash
curl -X DELETE "http://localhost:8080/api/genres/1?force=true"
```


## Project Structure

```text
movies-api/
├── cmd/
│   └── api/
│       └── main.go
├── data/
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── internal/
│   ├── database/
│   │   ├── db.go
│   │   └── movies.db
│   ├── handler/
│   │   ├── actor.go
│   │   ├── error_handler.go
│   │   ├── genre.go
│   │   └── movie.go
│   ├── models/
│   │   ├── errors/
│   │   │   └── errors.go
│   │   ├── models.go
│   │   └── query.go
│   ├── repository/
│   │   ├── actor.go
│   │   ├── actor_test.go
│   │   ├── genre.go
│   │   ├── genre_test.go
│   │   ├── movie.go
│   │   ├── movie_test.go
│   │   └── shared_test.go
│   └── service/
│       ├── actor.go
│       ├── actor_test.go
│       ├── genre.go
│       ├── movie.go
│       └── movie_test.go
├── migrations/
│   ├── 20260820220000_initial_schema.sql
│   └── 20260820230000_seed_data.sql
├── postman/
│   └── Movie Database API.postman_collection.json
├── static/
│   └── assets/
│       └── banner.png
└── README.md
```

## Team

| Name | ⬇ | ⬇ |
|---|---|---|
| Muhammad Owais Javed[Lead] | [GitHub](https://github.com/muhammad-owais-javed/) | [LinkedIn](https://www.linkedin.com/in/muhammad-owais-javed/) |
| Asha Gaire | [GitHub](https://github.com/ashagaire) | [LinkedIn](https://www.linkedin.com/in/asha-gaire-2b532217b/) |
| Syed Najam Ul Hassan Kazmi | [GitHub](https://github.com/Najam-Hassan-Kazmi) | [LinkedIn](https://www.linkedin.com/in/najam-ul-hassan-indie-web-developer/) |

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
