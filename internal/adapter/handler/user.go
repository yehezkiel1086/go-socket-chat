package handler

import (
	"go-socket/internal/core/domain"
	"go-socket/internal/core/port"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc port.UserService
}

func InitUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (uh *UserHandler) Register(c *gin.Context) {
	// bind input
	var input *RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username and password are required.",
		})
		return
	}

	// create new user
	_, err := uh.svc.CreateUser(c, &domain.User{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return		
	}

	// success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully.",
	})
}
