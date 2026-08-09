package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"movies-api/internal/database"
	"movies-api/internal/models"
	"movies-api/internal/repository"
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

	fmt.Printf(">> TEST: Genre Repository Call...\n")
	genreRepo := repository.NewGenreRepository(db)
	
	newGenre := &models.Genre{
		Name: "Horror",
	}

	fmt.Printf("\n--- Create Method ---\n")
	fmt.Printf("Initial Value > newGenre: %+v\n", newGenre)

	err = genreRepo.Create(newGenre)
	if err != nil {
		log.Printf("Error creating genre: %v\n", err)
	} else {
		fmt.Printf("Update Value > newGenre:  %+v\n", newGenre)
	}
	fmt.Printf("---------------------\n")
	
	fmt.Printf("\n--- GetAll Method ---\n")
	allGenres, err := genreRepo.GetAll()
	if err != nil {
		log.Printf("Error getting genre: %v\n", err)
	} else {
		fmt.Printf("Found %d genres: \n", len(allGenres))
		for _, g := range allGenres {
			fmt.Printf(" - ID: %d, Name: %s\n", g.ID, g.Name)
		}
	}
	fmt.Printf("---------------------\n")

	fmt.Printf("\n--- Update Method ---\n")
	genreToUpdate := &models.Genre{
		ID:   3,
		Name: "Horror",
	}
	err = genreRepo.Update(genreToUpdate)
	if err != nil {
		log.Printf("Error updating genre: %v\n", err)
	} else {
		fmt.Println("Successfully updated genre ID'!")
	}

	fmt.Printf("---------------------\n")

	fmt.Printf("\n--- Delete Method ---\n")
	err = genreRepo.Delete(3)
	if err != nil {
		log.Printf("Error deleting genre: %v\n", err)
	} else {
		fmt.Println("Successfully deleted genre ID!")
	}

	fmt.Printf("---------------------\n")

	fmt.Printf("\n--- Verification Method ---\n")
	genre, err := genreRepo.GetByID(2)
	if err != nil {
		fmt.Printf("Verification: %v\n", err) 
	} else {
		fmt.Printf(" - ID: %d, Name: %s\n", genre.ID, genre.Name)
	}
	fmt.Printf("---------------------\n")

	port := ":8080"
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", HealthCheck)

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
