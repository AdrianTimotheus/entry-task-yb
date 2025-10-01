package services

import (
	"awesomeProject/dto"
	"awesomeProject/models"
	"awesomeProject/repositories"
)

type ActorService interface {
	CreateActor(actor *dto.CreateActorReq) (*models.Actor, error)
	GetActor(id int) (*models.Actor, error)
	GetActors() ([]*models.Actor, error)
	UpdateActor(actor *dto.UpdateActorReq) (*models.Actor, error)
	DeleteActor(id int) (*models.Actor, error)
}

type actorService struct {
	actorRepository repositories.ActorRepository
}

type ActorServiceConfig struct {
	ActorRepository repositories.ActorRepository
}

func NewActorService(config *ActorServiceConfig) ActorService {
	return &actorService{
		actorRepository: config.ActorRepository,
	}
}

func (s *actorService) CreateActor(actor *dto.CreateActorReq) (*models.Actor, error) {
	newActor := &models.Actor{
		Name: actor.Name,
	}

	createdActor, err := s.actorRepository.CreateActor(newActor)

	if err != nil {
		return nil, err
	}

	return createdActor, nil
}

func (s *actorService) GetActor(id int) (*models.Actor, error) {
	actor, err := s.actorRepository.GetActor(id)
	if err != nil {
		return nil, err
	}
	return actor, nil
}

func (s *actorService) GetActors() ([]*models.Actor, error) {
	actors, err := s.actorRepository.GetActors()
	if err != nil {
		return nil, err
	}
	return actors, nil
}

func (s *actorService) UpdateActor(actor *dto.UpdateActorReq) (*models.Actor, error) {
	updatedActor := models.Actor{
		ID:   actor.ID,
		Name: actor.Name,
	}

	updateActor, err := s.actorRepository.UpdateActor(&updatedActor)
	if err != nil {
		return nil, err
	}
	return updateActor, nil
}

func (s *actorService) DeleteActor(id int) (*models.Actor, error) {
	actor, err := s.actorRepository.DeleteActor(id)
	if err != nil {
		return nil, err
	}
	return actor, nil
}
