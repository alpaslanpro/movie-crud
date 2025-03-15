package repositories

import (
	"github.com/alpaslanpro/movie-crud/models"
	"gorm.io/gorm"
)

type MovieRepository interface {
	Create(movie *models.Movie) (*models.Movie, error)
	FindByID(id uint) (*models.Movie, error)
	FindAll() ([]*models.Movie, error)
	FindWithPagination(page, pageSize int, filter string) ([]*models.Movie, error)
	Update(movie *models.Movie) (*models.Movie, error)
	Delete(id uint) error
}

type GormMovieRepository struct {
	db *gorm.DB
}

func NewGormMovieRepository(db *gorm.DB) *GormMovieRepository {
	return &GormMovieRepository{db: db}
}

func (r *GormMovieRepository) Create(movie *models.Movie) (*models.Movie, error) {
	return nil, nil
}

func (r *GormMovieRepository) FindByID(id uint) (*models.Movie, error) {
	return nil, nil
}

func (r *GormMovieRepository) FindAll() ([]*models.Movie, error) {
	return nil, nil
}

func (r *GormMovieRepository) FindWithPagination(page, pageSize int, filter string) ([]*models.Movie, error) {
	return nil, nil
}

func (r *GormMovieRepository) Update(movie *models.Movie) (*models.Movie, error) {
	return nil, nil
}

func (r *GormMovieRepository) Delete(id uint) error {
	return nil
}
