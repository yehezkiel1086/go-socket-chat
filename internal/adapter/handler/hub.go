package handler

import (
	"go-socket/internal/core/domain"
	"go-socket/internal/core/port"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HubHandler struct {
	svc port.HubService
}

func InitHubHandler(svc port.HubService) *HubHandler {
	return &HubHandler{
		svc: svc,
	}
}

type CreateRoomReq struct {
	ID string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (hh *HubHandler) CreateRoom(c *gin.Context) {
	// bind inputs
	var input *CreateRoomReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id and name are required",
		})
		return
	}

	// create room
	_, err := hh.svc.CreateRoom(c, &domain.Room{
		ID: input.ID,
		Name: input.Name,
		Clients: make(map[string]*domain.Client),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// return response
	c.JSON(http.StatusCreated, input)
}
