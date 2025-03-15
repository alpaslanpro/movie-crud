package main

import (
	"fmt"

	"github.com/alpaslanpro/movie-crud/db"
	"github.com/alpaslanpro/movie-crud/repositories"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		db.Module,
		repositories.Module,
		fx.Invoke(testDatabaseConnection),
	)

	app.Run()
}

func testDatabaseConnection(repo repositories.MovieRepository) {
	fmt.Println("Movie repository has been successfully injected!")
}
