package main

import (
	"fmt"
	"os"
	"log"
	"movies-api/internal/database"
)

func main(){

	fmt.Println("movies-api")

	err := os.MkdirAll("./data", 0755)
	if err != nil {
		log.Fatalf(">> FATAL ERROR: Failed to create directory \n%v", err)
	}

	dbPath := "./data/movies.db"
	db, err := database.New(dbPath)
	if err != nil {
		log.Fatalf(">> ERROR: %v", err)
	}

	defer db.Close()

	migrationPath := "./migrations/0000001_init_schema.up.sql"
	err = database.RunMigration(db, migrationPath)
	if err != nil {
		log.Fatalf(">> FATAL ERROR: Could not run migration \n%v", err)
	}

	fmt.Println(">> Database Run Complete!")
	
}