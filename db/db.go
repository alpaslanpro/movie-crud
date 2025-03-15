package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/alpaslanpro/movie-crud/models"
	"github.com/alpaslanpro/movie-crud/repositories"
	"go.uber.org/fx"
)

func NewPostgresDB() (*gorm.DB, error) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	dsn := databaseURL
	if dsn == "" {
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PASSWORD"),
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return nil, err
	}

	err = db.AutoMigrate(&models.Movie{}, &models.Actor{}, &models.Genre{}, &models.User{})
	if err != nil {
		fmt.Println("Failed to migrate database:", err)
		return nil, err
	}

	fmt.Println("Database connected successfully")
	return db, nil
}

func NewMovieRepository(db *gorm.DB) repositories.MovieRepository {
	return repositories.NewGormMovieRepository(db)
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return repositories.NewGormUserRepository(db)
}

var Module = fx.Options(
	fx.Provide(NewPostgresDB),
	fx.Provide(NewMovieRepository),
	fx.Provide(NewUserRepository),
)
