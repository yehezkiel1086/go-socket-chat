package handler

import (
	"go-socket/internal/core/domain"
	"go-socket/internal/core/port"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ClientHandler struct {
	svc port.ClientService
}

func InitClientHandler(svc port.ClientService) *ClientHandler {
	return &ClientHandler{
		svc: svc,
	}
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func (ch *ClientHandler) JoinRoom(c *gin.Context) {
	// get queries and params
	roomId := c.Param("room_id")
	username := c.Query("username")
	clientId := c.Query("client_id")
	if roomId == "" || username == "" || clientId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "room_id, username and client_id are required",
		})
		return
	}

	// create new connection
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// join room
	ch.svc.JoinRoom(c, &domain.Client{
		Conn: conn,
		RoomID: roomId,
		Username: username,
		ID: clientId,
		Message: make(chan *domain.Message, 10),
	})
}
