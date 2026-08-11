package repository

import (
	"movies-api/internal/models"
	"testing"
)

func TestGenreRepository_CRUD(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	repo := NewGenreRepository(db)

	genre := &models.Genre{Name: "Action"}
	err := repo.Create(genre)
	if err != nil || genre.ID == 0 {
		t.Fatalf("Failed to create genre: %v", err)
	}

	fetched, err := repo.GetByID(genre.ID)
	if err != nil || fetched.Name != "Action" {
		t.Errorf("Failed to fetch genre: %v", err)
	}

	repo.Create(&models.Genre{Name: "Comedy"})
	all, err := repo.GetAll()
	if err != nil || len(all) != 2 {
		t.Errorf("Expected 2 genres, got %d", len(all))
	}

	genre.Name = "Thriller"
	err = repo.Update(genre)
	if err != nil {
		t.Errorf("Failed to update genre: %v", err)
	}

	err = repo.Delete(genre.ID)
	if err != nil {
		t.Errorf("Failed to delete genre: %v", err)
	}
}
