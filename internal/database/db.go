package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func New(dbPath string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	fmt.Println(">> Connection to Database ...")
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println(">> PING: Successful!")
	fmt.Println(">> Connection to SQLite Established!")
	return db, nil
}

func RunMigration(db *sql.DB, migrationFilePath string) error {

	fmt.Println(">> Checking Database Schema ...")

	sqlBytes, err := os.ReadFile(migrationFilePath)
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	fmt.Println(">> Database Schema Check Complete!")

	return nil
}
