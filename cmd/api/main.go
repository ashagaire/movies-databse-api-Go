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
		log.Fatalf("Failed to create directory: %v", err)
	}

	dbPath := "./data/movies.db"
	db, err := database.New(dbPath)
	if err != nil {
		log.Fatalf(">> ERROR: %v", err)
	}

	defer db.Close()
	
}