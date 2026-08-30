package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Global db variable with an InitDB function\
// goose -dir migrations sqlite3 ./internal/db/test.db up

func InitDB() (*sql.DB, error){
	var err error
	db, err := sql.Open("sqlite3", "../../internal/db/test.db?_foreign_keys=on")
	if err != nil {
		log.Fatal(err)
		return nil, err

	}
	fmt.Println("database linked sucessfully")
	return db, nil
}

func seedData(db *sql.DB, seeDFilePath string) error {
	script, err := os.ReadFile(seeDFilePath)
	if err != nil {
		fmt.Println("error in finding file")
		return err
	}

	_, err = db.Exec(string(script))
	if err != nil {
		fmt.Println("error in seeding data ")
		return err
	}
	return nil
	
}