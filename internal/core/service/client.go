package service

import (
	"context"
	"fmt"
	"go-socket/internal/core/domain"
	"go-socket/internal/core/port"
)

type ClientService struct {
	hubRepo port.HubRepository
}

func InitClientService(hubRepo port.HubRepository) *ClientService {
	return &ClientService{
		hubRepo: hubRepo,
	}
}

func (cs *ClientService) JoinRoom(ctx context.Context, cl *domain.Client) {
	// create join message
	msg := &domain.Message{
		Content: fmt.Sprintf("%v has joined the chat", cl.Username),
		RoomID: cl.RoomID,
		Username: cl.Username,
	}

	hub := cs.hubRepo.GetHub(ctx)

	// register client
	hub.Register <- cl	

	// broadcast join message
	hub.Broadcast <- msg

	// write and read messages
	go cl.WriteMessage()
	cl.ReadMessage(hub)
}
