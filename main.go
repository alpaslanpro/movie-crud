package main

import (
	"fmt"

	"github.com/alpaslanpro/movie-crud/db"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func main() {
	app := fx.New(
		db.Module,
		fx.Invoke(testDatabaseConnection),
	)

	app.Run()
}

func testDatabaseConnection(dbInstance *gorm.DB) {
	sqlDB, err := dbInstance.DB()
	if err != nil {
		panic("Failed to retrieve database instance")
	}

	err = sqlDB.Ping()
	if err != nil {
		panic("Database connection failed")
	}

	fmt.Println("Successfully connected to the database!")
}
