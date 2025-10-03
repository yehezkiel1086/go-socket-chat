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

func (hr *HubRepository) GetRooms(ctx context.Context) ([]*domain.Room, error) {
	if len(hr.hub.Rooms) == 0 {
		return []*domain.Room{}, errors.New("no rooms found")
	}

	rooms := []*domain.Room{}
	for _, r := range hr.hub.Rooms {
		rooms = append(rooms, r)
	}

	return rooms, nil
}
