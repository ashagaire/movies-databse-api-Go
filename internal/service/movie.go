package service

import (
	"errors"
	"fmt"
	"strings"

	"movies-api/internal/models"
	"movies-api/internal/repository"
)

type MovieService struct {
	repo *repository.MovieRepository
}

func NewMovieService(repo *repository.MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) Create(movie *models.Movie) error {
	
	movie.Title = strings.TrimSpace(movie.Title)
	if movie.Title == "" {
		return errors.New("movie title cannot be empty")
	}

	if movie.ReleaseYear < 1888 {
		return errors.New("release year must be 1888 or later")
	}

	if movie.Duration <= 0 {
		return errors.New("duration must be greater than 0 minutes")
	}

	err := s.repo.Create(movie)
	if err != nil {
		return fmt.Errorf("failed to create movie: %w", err)
	}

	return nil
}

func (s *MovieService) GetAll() ([]models.Movie, error) {
	
	return s.repo.GetAll()

}