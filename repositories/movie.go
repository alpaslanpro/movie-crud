package repositories

import (
	"github.com/alpaslanpro/movie-crud/models"
	"gorm.io/gorm"
)

type MovieRepository interface {
	Create(movie *models.Movie) error
	FindByID(id uint) (*models.Movie, error)
	FindAll() ([]models.Movie, error)
	Update(movie *models.Movie) error
	Delete(id uint) error
}

type GormMovieRepository struct {
	db *gorm.DB
}

func NewGormMovieRepository(db *gorm.DB) *GormMovieRepository {
	return &GormMovieRepository{db: db}
}

func (r *GormMovieRepository) Create(movie *models.Movie) error {
	return r.db.Create(movie).Error
}

func (r *GormMovieRepository) FindByID(id uint) (*models.Movie, error) {
	var movie models.Movie
	err := r.db.First(&movie, id).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *GormMovieRepository) FindAll() ([]models.Movie, error) {
	var movies []models.Movie
	err := r.db.Find(&movies).Error
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *GormMovieRepository) Update(movie *models.Movie) error {
	return r.db.Save(movie).Error
}

func (r *GormMovieRepository) Delete(id uint) error {
	return r.db.Delete(&models.Movie{}, id).Error
}
