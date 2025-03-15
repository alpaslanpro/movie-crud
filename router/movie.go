package router

import (
	"net/http"
	"strconv"

	"github.com/alpaslanpro/movie-crud/models"
	"github.com/alpaslanpro/movie-crud/pkg"
	"github.com/alpaslanpro/movie-crud/repositories"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type Response struct {
	Message string `json:"message"`
}

type Movie struct {
	Title  string   `json:"title"`
	Genres []string `json:"genres" gorm:"type:text[]" swaggertype:"array,string"` // Define as an array of strings
	Actors []string `json:"actors" gorm:"type:text[]" swaggertype:"array,string"` // Define as an array of strings
}

type MovieHandler struct {
	MovieRepo repositories.MovieRepository
}

func NewMovieHandler(movieRepo repositories.MovieRepository) *MovieHandler {
	return &MovieHandler{MovieRepo: movieRepo}
}

func (h *MovieHandler) PostMovie(ctx *gin.Context) {
	var movie models.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pkg.Validate.Struct(movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if movie.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Movie title is required"})
		return
	}
	if len(movie.Genres) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "At least one genre is required"})
		return
	}
	if len(movie.Actors) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "At least one actor is required"})
		return
	}

	res, err := h.MovieRepo.Create(&movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func (h *MovieHandler) GetMovies(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("page_size", "10")
	filter := ctx.DefaultQuery("filter", "")

	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	res, err := h.MovieRepo.FindWithPagination(pageInt, pageSizeInt, filter)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *MovieHandler) GetMovie(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	res, err := h.MovieRepo.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *MovieHandler) UpdateMovie(ctx *gin.Context) {
	var movie models.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("id")
	movieID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}
	movie.ID = uint(movieID)
	res, err := h.MovieRepo.Update(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *MovieHandler) DeleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	movieID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}
	err = h.MovieRepo.Delete(uint(movieID))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "movie deleted successfully"})
}
