package repository

import (
	"fmt"
	"database/sql"
	"movies-api/internal/models"
)

type ActorRepository struct {
	db *sql.DB
}

func NewActorRepository(db *sql.DB) *ActorRepository {
	return &ActorRepository{
		db: db,
	}
}

func (r *ActorRepository) Create(actor *models.Actor) error {

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