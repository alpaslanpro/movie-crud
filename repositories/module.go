package repositories

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return NewGormMovieRepository(db)
}

var Module = fx.Provide(NewMovieRepository)
