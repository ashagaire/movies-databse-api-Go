package service

import (
	"testing"

	"movies-api/internal/models"
	"movies-api/internal/repository"
)

func TestMovieService_CreateValidation(t *testing.T) {

	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewMovieRepository(db)
	svc := NewMovieService(repo)

	tests := []struct {
		name          string
		movie         *models.Movie
		expectedError string
	}{
		{
			name: "Valid Movie",
			movie: &models.Movie{
				Title:       "The Matrix",
				ReleaseYear: 1999,
				Duration:    136,
			},
			expectedError: "",
		},
		{
			name: "Empty Title",
			movie: &models.Movie{
				Title:       "   ",
				ReleaseYear: 1999,
				Duration:    136,
			},
			expectedError: "movie title cannot be empty",
		},
		{
			name: "Year Too Old",
			movie: &models.Movie{
				Title:       "Old Movie",
				ReleaseYear: 1800, // Before 1888
				Duration:    90,
			},
			expectedError: "release year must be 1888 or later",
		},
		{
			name: "Invalid Duration",
			movie: &models.Movie{
				Title:       "Short Movie",
				ReleaseYear: 2000,
				Duration:    0, // Invalid!
			},
			expectedError: "duration must be greater than 0 minutes",
		},
	}

	for _, tc := range tests {
		
		t.Run(tc.name, func(t *testing.T) {
			
			err := svc.Create(tc.movie)

			if tc.expectedError != "" && err == nil {
				t.Errorf("Expected error '%s', but got nil", tc.expectedError)
			}

			if err != nil && err.Error() != tc.expectedError {
				t.Errorf("Expected error '%s', but got '%v'", tc.expectedError, err)
			}
		})
	}
}