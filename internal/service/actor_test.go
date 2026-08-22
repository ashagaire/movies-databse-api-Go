package service

import (
	"database/sql"
	"testing"
	_ "github.com/mattn/go-sqlite3"
	"movies-api/internal/models"
	"movies-api/internal/repository"
)

func setupTestDB(t *testing.T) *sql.DB {

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	schema := `
		CREATE TABLE actors (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, birth_date TEXT NOT NULL);
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