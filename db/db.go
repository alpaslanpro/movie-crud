package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		return nil, err
	}

	return db, nil
}

var Module = fx.Provide(NewPostgresDB)
