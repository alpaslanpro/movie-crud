package router

import (
	"fmt"

	"github.com/alpaslanpro/movie-crud/repositories"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type HandlerStore struct {
	UserHandler  *UserHandler
	MovieHandler *MovieHandler
}

func ProvideHandlerStore(store *repositories.Store) *HandlerStore {
	return &HandlerStore{
		UserHandler:  NewUserHandler(store.UserRepo),
		MovieHandler: NewMovieHandler(store.MovieRepo),
	}
}

func ProvideRouter(h *HandlerStore) *gin.Engine {
	r := gin.Default()

	fmt.Println("Router successfully created:", r)

	r.POST("/login", h.UserHandler.LoginHandler)
	r.POST("/register", h.UserHandler.RegisterHandler)

	r.GET("/movies", h.MovieHandler.GetMovies)
	r.GET("/movies/:id", h.MovieHandler.GetMovie)
	r.POST("/movies", h.MovieHandler.PostMovie)
	r.PUT("/movies/:id", h.MovieHandler.UpdateMovie)
	r.DELETE("/movies/:id", h.MovieHandler.DeleteMovie)

	return r
}

var Module = fx.Module("handlers",
	fx.Provide(ProvideHandlerStore),
	fx.Provide(ProvideRouter),
)
