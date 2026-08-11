package repository

import (
	"movies-api/internal/models"
	"testing"
)

func TestMovieRepository_CreateAndGet(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	
	genreRepo := NewGenreRepository(db)
	actorRepo := NewActorRepository(db)
	movieRepo := NewMovieRepository(db)

	genre := &models.Genre{Name: "Sci-Fi"}
	genreRepo.Create(genre)
	
	actor := &models.Actor{Name: "Keanu Reeves", BirthDate: "1964-09-02"}
	actorRepo.Create(actor)

	movie := &models.Movie{
		Title:       "The Matrix",
		ReleaseYear: 1999,
		Duration:    136,
		Genres:      []models.Genre{*genre},
		Actors:      []models.Actor{*actor},
	}

	err := movieRepo.Create(movie)
	if err != nil {
		t.Fatalf("Failed to create movie: %v", err)
	}

	fetched, err := movieRepo.GetByID(movie.ID)
	if err != nil {
		t.Fatalf("Failed to fetch movie: %v", err)
	}

	if fetched.Title != "The Matrix" {
		t.Errorf("Expected 'The Matrix', got '%s'", fetched.Title)
	}
	if len(fetched.Genres) != 1 || fetched.Genres[0].Name != "Sci-Fi" {
		t.Errorf("Failed to fetch movie genres correctly")
	}
	if len(fetched.Actors) != 1 || fetched.Actors[0].Name != "Keanu Reeves" {
		t.Errorf("Failed to fetch movie actors correctly")
	}
}
