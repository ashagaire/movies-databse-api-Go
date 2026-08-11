package repository

import (
	"testing"
	"movies-api/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

// // setupTestDB creates in-memory database for our tests.
// func setupTestDB(t *testing.T) *sql.DB {
	
// 	// ":memory:" To create a temporary database in RAM
// 	db, err := sql.Open("sqlite3", ":memory:")
// 	if err != nil {
// 		t.Fatalf("Failed to open test database: %v", err)
// 	}

// 	schema := `
// 	CREATE TABLE actors (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		name TEXT NOT NULL,
// 		birth_date TEXT NOT NULL
// 	);`
	
// 	_, err = db.Exec(schema)
// 	if err != nil {
// 		t.Fatalf("Failed to create test schema: %v", err)
// 	}

// 	return db
// }


func TestActorRepository_Create(t *testing.T) {
	
	db := setupTestDB(t)
	defer db.Close()

	repo := NewActorRepository(db)

	actor := &models.Actor{
		Name:      "Keanu Reeves",
		BirthDate: "1964-09-02",
	}

	err := repo.Create(actor)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if actor.ID == 0 {
		t.Errorf("Expected actor ID to be set, got 0")
	}
}

func TestActorRepository_GetByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	repo := NewActorRepository(db)

	repo.Create(&models.Actor{Name: "Tom Hanks", BirthDate: "1956-07-09"})

	fetchedActor, err := repo.GetByID(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if fetchedActor.Name != "Tom Hanks" {
		t.Errorf("Expected name 'Tom Hanks', got '%s'", fetchedActor.Name)
	}
}

func TestActorRepository_GetAll(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	repo := NewActorRepository(db)

	repo.Create(&models.Actor{Name: "Actor 1", BirthDate: "2000-01-01"})
	repo.Create(&models.Actor{Name: "Actor 2", BirthDate: "2000-01-01"})

	actors, err := repo.GetAll()
	if err != nil || len(actors) != 2 {
		t.Errorf("Expected 2 actors, got %d. Error: %v", len(actors), err)
	}
}

func TestActorRepository_UpdateAndDelete(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	repo := NewActorRepository(db)

	actor := &models.Actor{Name: "Old Name", BirthDate: "2000-01-01"}
	repo.Create(actor)

	actor.Name = "New Name"
	err := repo.Update(actor)
	if err != nil {
		t.Errorf("Failed to update actor: %v", err)
	}

	fetched, _ := repo.GetByID(actor.ID)
	if fetched.Name != "New Name" {
		t.Errorf("Expected 'New Name', got '%s'", fetched.Name)
	}

	err = repo.Delete(actor.ID)
	if err != nil {
		t.Errorf("Failed to delete actor: %v", err)
	}

	_, err = repo.GetByID(actor.ID)
	if err == nil {
		t.Errorf("Expected error when fetching deleted actor, got nil")
	}
}