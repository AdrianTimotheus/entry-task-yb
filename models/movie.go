package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Genre       string `json:"genre"`
}
