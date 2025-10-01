package repositories

import (
	"awesomeProject/models"
	"gorm.io/gorm"
)

type ActorRepository interface {
	CreateActor(actor *models.Actor) (*models.Actor, error)
	GetActor(id int) (*models.Actor, error)
	GetActors() ([]*models.Actor, error)
	UpdateActor(actor *models.Actor) (*models.Actor, error)
	DeleteActor(id int) (*models.Actor, error)
}

type actorRepository struct {
	db *gorm.DB
}

type ActorRepositoryConfig struct {
	DB *gorm.DB
}

func NewActorRepository(config *ActorRepositoryConfig) ActorRepository {
	return &actorRepository{db: config.DB}
}

func (r *actorRepository) CreateActor(actor *models.Actor) (*models.Actor, error) {
	result := r.db.Create(actor)
	if result.Error != nil {
		return nil, result.Error
	}
	return actor, nil
}

func (r *actorRepository) GetActor(id int) (*models.Actor, error) {
	var actor models.Actor
	result := r.db.Where("id = ?", id).First(&actor)
	if result.Error != nil {
		return nil, result.Error
	}
	return &actor, nil
}

func (r *actorRepository) GetActors() ([]*models.Actor, error) {
	var actors []*models.Actor
	result := r.db.Find(&actors)
	if result.Error != nil {
		return nil, result.Error
	}
	return actors, nil
}

func (r *actorRepository) UpdateActor(actor *models.Actor) (*models.Actor, error) {
	result := r.db.Save(&actor)
	if result.Error != nil {
		return nil, result.Error
	}
	return actor, nil
}

func (r *actorRepository) DeleteActor(id int) (*models.Actor, error) {
	var actor models.Actor
	result := r.db.Where("id = ?", id).First(&actor)
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.db.Where("id = ?", id).Delete(&models.Actor{})
	if result.Error != nil {
		return nil, result.Error
	}

	return &actor, nil
}
