package main

import (
	"fmt"
	"log"
	internal "movies-api/internal/db"
)

func main(){

	fmt.Println("movies-api")
	db, err:= internal.InitDB()
	if err != nil {
		log.Println("Error in database initialization.")
        log.Fatal(err)
    }
	defer db.Close()
}