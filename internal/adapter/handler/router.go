package handler

import (
	"go-socket/internal/adapter/config"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func InitRouter(
	config *config.HTTP,
	userHandler UserHandler,
	authHandler AuthHandler,
	hubHandler HubHandler,
	clientHandler ClientHandler,
) (*Router, error) {
	r := gin.New()

	v1 := r.Group("/api/v1") // public routes
	v1.POST("/register", userHandler.Register)
	v1.POST("/login", authHandler.Login)

	// web socket routes
	ws := v1.Group("/ws")
	ws.POST("/rooms", hubHandler.CreateRoom)
	ws.GET("/rooms", hubHandler.GetRooms)
	ws.GET("/rooms/join/:room_id", clientHandler.JoinRoom)

	return &Router{r}, nil
}

func (r *Router) Serve(uri string) error {
	return r.Run(uri)
}
