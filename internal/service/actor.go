package service

import (
	"fmt"
	"open-movies-db/internal/models"
	"open-movies-db/internal/repository"
)

type ActorService struct{
	repo *repository.ActorRepository
}

func NewActorService(repo *repository.ActorRepository) *ActorService{
	return &ActorService{
		repo: repo,
	}
}

func (s *ActorService) Create(actor *models.Actor) error{

	actor.Name = strings.TrimSpace(actor.Name)
	if actor.Name == "" {
		return errors.New("actor name cannot be empty")
	}

	_, err := time.Parse("2006-01-02", actor.BirthDate)
	if err != nil {
		return errors.New("birth date must be in YYYY-MM-DD format")
	}

	err := s.repo.Create(actor)
	if err != nil {
		return fmt.Errorf("failed to create actor: %w", err)
	}

	return nil

}

func (s *ActorService) GetAll() ([]models.Actor, error) {
	return s.repo.GetAll()
}