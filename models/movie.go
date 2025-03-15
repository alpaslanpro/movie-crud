package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Director    string `gorm:"not null"`
	Year        int    `gorm:"not null"`
}