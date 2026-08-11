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

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer tx.Rollback()

	
	query := `INSET INTO movies (title, release_year, duration) VALUES (?,?,?)`

	result, err := tx.Exec(query, movie.Title, movie.ReleaseYear, movie.Duration)
	if err != nil {
		return fmt.Errorf("failed to enter movie: %w", err)
	}

	movieID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get movieID: %w", err)
	}

	movie.ID = movieID

	for _, genre := range movie.Genres(
		
		genreQuery := `INSERT INTO movie_genres (movie_id, genre_id) VALUES (?, ?)`
		_, err := tx.Exec(genreQuery, movie.ID, genre.ID)
		if err != nil{
			return fmt.Errorf("failed to link genre ID %d: %w", genre.ID, err)
		}

	)

	for _, actor := range movie.Actors(

		actorQuery := `INSERT INTO movie_actors (movie_id, actor_id) VALUES (?, ?)`
		_, err := tx.Exec(actorQuery, movie.ID, actor.ID)
		if err != nil{
			return fmt.Errorf("failed to link actor ID %d: %w", actor.ID, err)
		}

	)

	return nil

}