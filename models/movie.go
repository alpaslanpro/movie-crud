package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model 
	Title  string  `json:"title"`
	Year   string  `json:"year"`
	Genres []Genre `json:"genres" gorm:"many2many:movie_genres;"`
	Actors []Actor `json:"actors" gorm:"many2many:movie_actors;"`
}

type Genre struct {
	gorm.Model
	Name string `json:"name"`
}

type Actor struct {
	gorm.Model
	Name string `json:"name"`
}
