package repositories

import (
	"errors"

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
	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	// Check and link existing actors
	var existingActors []models.Actor
	for _, actor := range movie.Actors {
		var existingActor models.Actor
		if err := tx.Where("name = ?", actor.Name).First(&existingActor).Error; err == nil {
			existingActors = append(existingActors, existingActor)
		} else {
			existingActors = append(existingActors, actor) // New actor
		}
	}
	movie.Actors = existingActors

	// Check and link existing genres
	var existingGenres []models.Genre
	for _, genre := range movie.Genres {
		var existingGenre models.Genre
		if err := tx.Where("name = ?", genre.Name).First(&existingGenre).Error; err == nil {
			existingGenres = append(existingGenres, existingGenre)
		} else {
			existingGenres = append(existingGenres, genre) // New genre
		}
	}
	movie.Genres = existingGenres

	// Create movie with associated data
	if err := tx.Create(movie).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return movie, nil
}


func (r *GormMovieRepository) FindByID(id uint) (*models.Movie, error) {
	var movie models.Movie
	if err := r.db.Preload("Genres").Preload("Actors").First(&movie, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("movie not found")
		}
		return nil, err
	}
	return &movie, nil
}

func (r *GormMovieRepository) FindAll() ([]*models.Movie, error) {
	var movies []*models.Movie
	if err := r.db.Preload("Genres").Preload("Actors").Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *GormMovieRepository) FindWithPagination(page, pageSize int, filter string) ([]*models.Movie, error) {
	var movies []*models.Movie
	offset := (page - 1) * pageSize
	query := r.db.Preload("Genres").Preload("Actors").Offset(offset).Limit(pageSize)

	if filter != "" {
		query = query.Where("title LIKE ?", "%"+filter+"%")
	}

	if err := query.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *GormMovieRepository) Update(movie *models.Movie) (*models.Movie, error) {
	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	// Check and link existing actors
	var existingActors []models.Actor
	for _, actor := range movie.Actors {
		var existingActor models.Actor
		if err := tx.Where("name = ?", actor.Name).First(&existingActor).Error; err == nil {
			existingActors = append(existingActors, existingActor)
		} else {
			existingActors = append(existingActors, actor) // New actor
		}
	}
	movie.Actors = existingActors

	// Check and link existing genres
	var existingGenres []models.Genre
	for _, genre := range movie.Genres {
		var existingGenre models.Genre
		if err := tx.Where("name = ?", genre.Name).First(&existingGenre).Error; err == nil {
			existingGenres = append(existingGenres, existingGenre)
		} else {
			existingGenres = append(existingGenres, genre) // New genre
		}
	}
	movie.Genres = existingGenres

	// Save movie with associations
	if err := tx.Model(&models.Movie{}).Where("id = ?", movie.ID).Updates(movie).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update associations
	if err := tx.Model(&movie).Association("Actors").Replace(movie.Actors); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Model(&movie).Association("Genres").Replace(movie.Genres); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return movie, nil
}

func (r *GormMovieRepository) Delete(id uint) error {
	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(&models.Movie{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
