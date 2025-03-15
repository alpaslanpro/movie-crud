package router

import (
	"github.com/alpaslanpro/movie-crud/repositories"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserRepo repositories.UserRepository
}

func NewUserHandler(userRepo repositories.UserRepository) *UserHandler {
	return &UserHandler{UserRepo: userRepo}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) LoginHandler(c *gin.Context) {

}

func (h *UserHandler) RegisterHandler(c *gin.Context) {

}
