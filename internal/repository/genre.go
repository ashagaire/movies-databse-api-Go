package repository

import (
	"database/sql"
	"errors"
	"fmt"
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
		return fmt.Errorf("failed to enter genre: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get new ID: %w", err)
	}

	genre.ID = id

	return nil
}

func (r *GenreRepository) GetAll() ([]models.Genre, error) {

	query := `SELECT id, name FROM genres`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query genres: %w", err)
	}

	defer rows.Close()

	var genres []models.Genre

	for rows.Next() {

		var g models.Genre

		err := rows.Scan(&g.ID, &g.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan genres: %w", err)
		}

		genres = append(genres, g)

	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("row interaction error: %w", err)
	}

	return genres, nil

}

func (r *GenreRepository) GetByID(id int64) (*models.Genre, error) {

	query := `SELECT id, name FROM genres WHERE id = ?`

	var g models.Genre

	err := r.db.QueryRow(query, id).Scan(&g.ID, &g.Name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("genre with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to query genre: %w", err)
	}

	return &g, nil
}

func (r *GenreRepository) Update(genre *models.Genre) error {
	query := `UPDATE genres SET name = ? WHERE id = ?`

	_, err := r.db.Exec(query, genre.Name, genre.ID)
	if err != nil {
		return fmt.Errorf("failed to update genre: %w", err)
	}

	return nil
}

func (r *GenreRepository) Delete(id int64) error {

	query := `DELETE FROM genres WHERE id = ?`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete genre: %w", err)
	}

	return nil

}

func (r *GenreRepository) ForceDelete(id int64) error {

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM movie_genres WHERE genre_id = ?`, id)
	if err != nil {
		return fmt.Errorf("failed to delete genre relationships: %w", err)
	}

	_, err = tx.Exec(`DELETE FROM genres WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("failed to delete genre: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil

}
