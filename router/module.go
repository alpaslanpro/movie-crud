package router

import (
	"fmt"

	"github.com/alpaslanpro/movie-crud/pkg/auth"
	"github.com/alpaslanpro/movie-crud/repositories"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/login", h.UserHandler.LoginHandler)
	r.POST("/register", h.UserHandler.RegisterHandler)

	r.GET("/movies", auth.AuthMiddleware(), h.MovieHandler.GetMovies)
	r.GET("/movies/:id", auth.AuthMiddleware(), h.MovieHandler.GetMovie)
	r.POST("/movies", auth.AuthMiddleware(), h.MovieHandler.PostMovie)
	r.PUT("/movies/:id", auth.AuthMiddleware(), h.MovieHandler.UpdateMovie)
	r.DELETE("/movies/:id", auth.AuthMiddleware(), h.MovieHandler.DeleteMovie)

	return r
}

var Module = fx.Module("handlers",
	fx.Provide(ProvideHandlerStore),
	fx.Provide(ProvideRouter),
)
