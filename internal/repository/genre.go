package repository

import (
	"fmt"
	"database/sql"
	"movies-api/internal/models"
)

type GenreRepository struct {
	db *sql.DB
}

func NewGenreRepository(db *sql.DB) *GenreRepository {
	return &GenreRepository{
		db: db,
	}
}

func (r *GenreRepository) Create(genre *models.Genre) error {

	query := `INSERT INTO genres (name) VALUES (?)`

	result, err := r.db.Exec(query, genre.Name)
	if err != nil {
		fmt.Errorf("failed to enter genre: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Errorf("failed to get new ID: %w", err)
	}

	genre.ID = id
	

	return nil
}