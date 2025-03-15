package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/alpaslanpro/movie-crud/db"
	"github.com/alpaslanpro/movie-crud/repositories"
	"github.com/alpaslanpro/movie-crud/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		db.Module,
		repositories.Module,
		router.Module,
		fx.Invoke(GinServer),
		fx.Invoke(func() { fmt.Println("Application has been invoked") }),
	)

	app.Run()
}

func GinServer(lc fx.Lifecycle, r *gin.Engine) *http.Server {

	fmt.Println("GinServer() function called")

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Gin server starting on :8080")
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					fmt.Printf("Error starting server: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping Gin server...")
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
