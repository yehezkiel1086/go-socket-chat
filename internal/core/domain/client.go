package domain

import "golang.org/x/net/websocket"

type Client struct {
	Conn     *websocket.Conn
	ID       string
	Username string
	RoomID   string
	Message chan *Message
}

type Message struct {
	Content string
	RoomID string
	Username string
}
