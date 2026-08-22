package service

import (
	"fmt"
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

}
