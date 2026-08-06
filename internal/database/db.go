package database

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func New(dbPath string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println(">> PING: Successful!*")
	fmt.Println("*** Connection Established to SQLite Database! ***")
	return db, nil
}