package main

import (
	"fmt"
	"log"
	"movies-api/internal/database"
	"net/http"
	"os"
)

func main() {

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

	port := ":8080"
	mux := http.NewServeMux()

	mux.HandleFunc("/health", HealthCheck)

	fmt.Printf(">> Starting Server ...\n")
	fmt.Printf(">> URL: http://localhost%s\n", port)

	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf(">> FATAL: Server Failed to start! \n%v", err)
	}

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(">> OK: API Health Check!"))
	fmt.Printf(">> OK: API Health Check!\n")
}
