package internal

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func BuildDb(){
	db, err := sql.Open("sqlite3", "../../internal/db/test.db?_foreign_keys=on")
	if err != nil {
		log.Fatal(err)

	}
	defer db.Close()
	
	
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database built sucessfully")

	//seeding demo data
	seedFile := "seed.sql"
	err = seedData(db, seedFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database seeded sucessfully")
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