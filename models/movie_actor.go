package models

import "gorm.io/gorm"

type Role int

const (
	Primary Role = iota + 1
	Supporting
)

type MovieActor struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primaryKey"`
	ActorID    uint   `json:"actor_id"`
	MovieID    uint   `json:"movie_id"`
	StarringAs string `json:"starring_as"`
	Role       Role   `json:"role"`
}
