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

	for _, genre := range movie.Genres{
		
		genreQuery := `INSERT INTO movie_genres (movie_id, genre_id) VALUES (?, ?)`
		_, err := tx.Exec(genreQuery, movie.ID, genre.ID)
		if err != nil{
			return fmt.Errorf("failed to link genre ID %d: %w", genre.ID, err)
		}

	}

	for _, actor := range movie.Actors{

		actorQuery := `INSERT INTO movie_actors (movie_id, actor_id) VALUES (?, ?)`
		_, err := tx.Exec(actorQuery, movie.ID, actor.ID)
		if err != nil{
			return fmt.Errorf("failed to link actor ID %d: %w", actor.ID, err)
		}

	}

	return nil

}


func (r *MovieRepository) GetAll() ([]models.Movie, error) {

	movieQuery := `SELECT id, title, release_year, duration FROM movies`
	movieRows, err := r.db.Query(movieQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query movies: %w", err)
	}
	
	defer movieRows.Close()

	movieMap := make(map[int64]*models.Movie)

	var movieIDs []int64

	for movieRows.Next(){
		
		var m models.Movie

		err := movieRows.Scan(&m.ID, &m.Title, &m.ReleaseYear, &m.Duration)
		if  err != nil {
			return nil, fmt.Errorf("failed to scan movie: %w", err)
		}

		m.Genres = []models.Genre{}
		m.Actors = []models.Actor{}

		movieMap[m.ID] = &m
		movieIDs = append(movieIDs, m.ID)

	}

	if len(movieMap) == 0 {
		return []models.Movie{}, nil
	}

	genreQuery := `

	`
	genreRows, err := r.db.Query(genreQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query movie genres: %w", err)
	}
	defer genreRows.Close()

	for genreRows.Next() {
		
		var movieID int64
		var g models.Genre

		err := genreRows.Scan(&movieID, &g.ID, &g.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan movie genre: %w", err)
		}

		movie, exists := movieMap[movieID]
		if exists == true {
			movie.Genres = append(movie.Genres, g)
		}
	}

	actorQuery := `

	`
	actorRows, err := r.db.Query(actorQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query movie actors: %w", err)
	}
	defer actorRows.Close()

	for actorRows.Next() {

		var movieID int64
		var a models.Actor
		
		err := actorRows.Scan(&movieID, &a.ID, &a.Name, &a.BirthDate)
		if  err != nil {
			return nil, fmt.Errorf("failed to scan movie actor: %w", err)
		}

		movie, exists := movieMap[movieID]
		if exists == true {
			movie.Actors = append(movie.Actors, a)
		}
	}

	var finalMovies []models.Movies
	
	for _, id := range movieIDs {
		finalMovies = append(finalMovies, *movieMap[id])
	}

	return finalMovies, nil
}