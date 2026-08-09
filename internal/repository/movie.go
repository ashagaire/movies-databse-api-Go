package repository

import (
	"fmt"
	"database/sql"
	"movies-api/internal/models"
)

type MovieRepository struct {

	db *sql.DB
}

func NewMovieRepository (db *sql.DB) *MovieRepository {

	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) Create(movie *models.Movie) error {

	query := `INSET INTO movies (title, release_year, duration) VALUES (?,?,?)`

	result, err := r.db.Exec(query, movie.Title, movie.ReleaseYear, movie.Duration)
	if err != nil {
		return fmt.Errorf("failed to enter movie: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get id: %w", err)
	}

	movie.ID = id

	return nil

}