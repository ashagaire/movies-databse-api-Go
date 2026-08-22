package service

import (
	"database/sql"
	"testing"
	"github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	schema := `
		CREATE TABLE actors (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, birth_date TEXT NOT NULL);
	`
	_, err = db.Exec(schema)
	if err != nil {
		t.Fatalf("Failed to create test schema: %v", err)
	}

	return db

}

func TestActorService_CreateValidation(){

}