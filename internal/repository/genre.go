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

func (r *GenreRepository) GetAll() ([]models.Genre, err){

	query := `SELECT id, name FROM genres`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query genres: %w", err)
	}

	defer row.Close()

	var genres []models.Genre

	for row.Next(){

		var g models.Genre

		err := rows.Scan(&g.ID, &g.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan genres: %w", err)
		}

		genres = append(genres, g)

	}

	err := rows.Err()
	if err != nil {
		return nil, fmt.Errorf("row interaction error: %w", err)
	}

	return genres, nil

}