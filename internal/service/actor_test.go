package service

import (
	"database/sql"
	"testing"
	_ "github.com/mattn/go-sqlite3"
	"movies-api/internal/models"
	"movies-api/internal/repository"
)

func setupTestDB(t *testing.T) *sql.DB {

	db, err := sql.Open("sqlite3", ":memory:?_foreign_keys=on")
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	schema := `
		CREATE TABLE IF NOT EXISTS movies (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			release_year INTEGER NOT NULL,
			duration INTEGER NOT NULL
		);
		CREATE TABLE actors (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			name TEXT NOT NULL, 
			birth_date TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS movie_actors (    
			movie_id INTEGER NOT NULL,
			actor_id INTEGER NOT NULL,
			PRIMARY KEY (movie_id, actor_id),
			FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE RESTRICT,
			FOREIGN KEY (actor_id) REFERENCES actors(id) ON DELETE RESTRICT
	);
	`
	_, err = db.Exec(schema)
	if err != nil {
		t.Fatalf("Failed to create test schema: %v", err)
	}

	return db

}

func TestActorService_CreateValidation(t *testing.T) {

	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewActorRepository(db)
	svc := NewActorService(repo)

	// Table Driven tsting


	/*
		type actorTestCase struct {
			name          string
			actor         *models.Actor
			expectedError string
		}
		tests := []actorTestCase{
			{
				name:          "Valid Actor",
				actor:         &models.Actor{Name: "Tom Cruise", BirthDate: "1962-07-03"},
				expectedError: "",
			},
		}
	*/

	tests := []struct {
		name          string
		actor         *models.Actor
		expectedError string	
		}{
		{
			name:          "Valid Actor",
			actor:         &models.Actor{Name: "Tom Cruise", BirthDate: "1962-07-03"},
			expectedError: "", // We expect NO error here
		},
		{
			name:          "Empty Name",
			actor:         &models.Actor{Name: "   ", BirthDate: "1962-07-03"},
			expectedError: "actor name cannot be empty",
		},
		{
			name:          "Invalid Date Format",
			actor:         &models.Actor{Name: "Tom Cruise", BirthDate: "07-03-1962"}, // Wrong format!
			expectedError: "birth date must be in YYYY-MM-DD format",
		},
		{
			name:          "Nonsense Date",
			actor:         &models.Actor{Name: "Tom Cruise", BirthDate: "not-a-date"},
			expectedError: "birth date must be in YYYY-MM-DD format",
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			
			err := svc.Create(tc.actor)

			// Expected an error but if doesnt get one
			if tc.expectedError != "" && err == nil {
				t.Errorf("Expected error '%s', but got nil", tc.expectedError)
			}

			// Got an error, but it wasn't expected one
			if err != nil && err.Error() != tc.expectedError {
				t.Errorf("Expected error '%s', but got '%v'", tc.expectedError, err)
			}
		})
	}

}


func TestActorService_DeleteLogic(t *testing.T) {

	db := setupTestDB(t)
	defer db.Close()
	
	repo := repository.NewActorRepository(db)
	svc := NewActorService(repo)

	actor := &models.Actor{Name: "Tom Hanks", BirthDate: "1956-07-09"}
	svc.Create(actor)


	db.Exec(`INSERT INTO movies (id, title, release_year, duration) VALUES (1, 'Cast Away', 2000, 143)`)
	db.Exec(`INSERT INTO movie_actors (movie_id, actor_id) VALUES (1, ?)`, actor.ID)

	err := svc.Delete(actor.ID, false)
	if err == nil {
		t.Errorf("Expected normal delete to fail because actor is linked to a movie, but it succeeded!")
	}

	err = svc.Delete(actor.ID, true)
	if err != nil {
		t.Errorf("Expected force delete to succeed, but got error: %v", err)
	}

	_, err = svc.GetByID(actor.ID)
	if err == nil {
		t.Errorf("Actor should be deleted, but was still found!")
	}
}
