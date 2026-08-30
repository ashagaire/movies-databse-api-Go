package main

import (
	"fmt"
	"log"
	database "movies-api/internal/database"
)

func main(){

	fmt.Println("movies-api")
	db, err:= database.InitDB()
	if err != nil {
		log.Println("Error in database initialization.")
        log.Fatal(err)
    }
	defer db.Close()
}