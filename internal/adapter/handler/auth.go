package handler

import (
	"go-socket/internal/core/domain"
	"go-socket/internal/core/port"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc port.AuthService
}

func InitAuthHandler(svc port.AuthService) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}

func (ah *AuthHandler) Login(c *gin.Context) {
	// input binding
	var input *domain.AuthRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username and password are required.",
		})
		return
	}

	// login
	res, err := ah.svc.Login(c, input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return		
	}

	// return response
	c.JSON(http.StatusOK, res)
}
