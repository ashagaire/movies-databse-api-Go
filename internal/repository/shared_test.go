package repository

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	schema := `
	CREATE TABLE genres (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE);
	CREATE TABLE actors (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, birth_date TEXT NOT NULL);
	CREATE TABLE movies (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, release_year INTEGER NOT NULL, duration INTEGER NOT NULL);
	CREATE TABLE movie_genres (movie_id INTEGER, genre_id INTEGER, PRIMARY KEY (movie_id, genre_id), FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE RESTRICT, FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE RESTRICT);
	CREATE TABLE movie_actors (movie_id INTEGER, actor_id INTEGER, PRIMARY KEY (movie_id, actor_id), FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE RESTRICT, FOREIGN KEY (actor_id) REFERENCES actors(id) ON DELETE RESTRICT);
	`
	
	_, err = db.Exec(schema)
	if err != nil {
		t.Fatalf("Failed to create test schema: %v", err)
	}

	return db
}
