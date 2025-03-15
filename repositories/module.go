package repositories

import (
  "go.uber.org/fx"
  "gorm.io/gorm"
)

// Repository dependencies
type Store struct {
  UserRepo  UserRepository
  MovieRepo MovieRepository
}

// ProvideStore initializes repositories with the database
func ProvideStore(db *gorm.DB) *Store {
  return &Store{
    UserRepo:  NewGormUserRepository(db),
    MovieRepo: NewGormMovieRepository(db),
  }
}

// Module exports repository dependencies
var Module = fx.Module("repositories",
  fx.Provide(ProvideStore),
)