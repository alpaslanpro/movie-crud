package repositories

import (
	"github.com/alpaslanpro/movie-crud/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{DB: db}
}

func (r *GormUserRepository) CreateUser(user *models.User) (*models.User, error) {
	return nil, nil
}

func (r *GormUserRepository) GetUserByUsername(username string) (*models.User, error) {
	return nil, nil
}
