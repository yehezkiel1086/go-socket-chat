package repository

import (
	"context"
	"errors"
	"go-socket/internal/core/domain"
)

type HubRepository struct {
	hub *domain.Hub
}

func InitHubRepository() *HubRepository {
	return &HubRepository{
		hub: &domain.Hub{
			Rooms: make(map[string]*domain.Room),
			Register: make(chan *domain.Client),
			Unregister: make(chan *domain.Client),
			Broadcast: make(chan *domain.Message, 5),
		},
	}
}

func (hr *HubRepository) CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error) {
	if _, ok := hr.hub.Rooms[room.ID]; ok {
		return &domain.Room{}, errors.New("room already exists")
	}

	hr.hub.Rooms[room.ID] = room

	return room, nil
}
