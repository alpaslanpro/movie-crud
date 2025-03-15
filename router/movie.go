package router

import (
	"github.com/alpaslanpro/movie-crud/repositories"
	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	MovieRepo repositories.MovieRepository
}

func NewMovieHandler(movieRepo repositories.MovieRepository) *MovieHandler {
	return &MovieHandler{MovieRepo: movieRepo}
}

func (h *MovieHandler) PostMovie(ctx *gin.Context) {

}

func (h *MovieHandler) GetMovies(ctx *gin.Context) {

}

func (h *MovieHandler) GetMovie(ctx *gin.Context) {

}

func (h *MovieHandler) UpdateMovie(ctx *gin.Context) {

}

func (h *MovieHandler) DeleteMovie(ctx *gin.Context) {

}
