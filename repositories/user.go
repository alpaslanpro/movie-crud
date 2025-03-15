package repositories

import (
    "github.com/alpaslanpro/movie-crud/models"
    "gorm.io/gorm"
    "errors"
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
    result := r.DB.Create(user)
    if result.Error != nil {
        return nil, result.Error
    }
    return user, nil
}

func (r *GormUserRepository) GetUserByUsername(username string) (*models.User, error) {
    var user models.User
    err := r.DB.Where("username = ?", username).First(&user).Error
    if err != nil {
      return nil, errors.New("user not found")
    }
    return &user, nil
  }
