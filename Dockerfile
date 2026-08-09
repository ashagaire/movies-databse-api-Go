FROM golang:1.26-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o movies-api ./cmd/api

FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates sqlite3 && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/movies-api .

COPY --from=builder /app/migrations ./migrations

RUN mkdir -p /app/data

EXPOSE 8080

CMD ["./movies-api"]
