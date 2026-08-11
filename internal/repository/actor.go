package repository

import (
	"fmt"
	"database/sql"
	"movies-api/internal/models"
)

type ActorRepository struct {
	db *sql.DB
}

func NewActorRepository(db *sql.DB) *ActorRepository{
	return &ActorRepository{
		db: db,
	}
}

func (r *ActorRepository) Create(actor *models.Actor) error{

	query := `INSERT INTO actors (name, birth_date) VALUES (?, ?)`

	result, err := r.db.Exec(query, actor.Name, actor.BirthDate)
	if err != nil {
		return fmt.Errorf("failed to enter actor: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get new ID: %w", err)
	}

	actor.ID = id

	return nil
}

func (r *ActorRepository) GetAll()([]models.Actor, error){

	query := `SELECT id, name, birth_date FROM actors`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query actors: %w", err)
	}

	defer rows.Close()

	var actors []models.Actor
	
	for rows.Next(){
		
		var a models.Actor

		err := rows.Scan(&a.ID, &a.Name, &a.BirthDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan actor: %w", err)
		}

		rows = append(actors, a)
	}
	
	err = rows.Err();
	if err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return actors, nil
}