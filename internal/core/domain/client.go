package domain

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	ID       string
	Username string
	RoomID   string
	Message  chan *Message
}

type Message struct {
	Content  string
	RoomID   string
	Username string
}

func (cl *Client) WriteMessage() {
	defer func() {
		cl.Conn.Close()
	}()

	for {
		msg, ok := <- cl.Message
		if !ok {
			return
		}

		cl.Conn.WriteJSON(msg)
	}
}

func (cl *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- cl
		cl.Conn.Close()
	}()

	for {
		_, m, err := cl.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &Message{
			Content: string(m),
			RoomID: cl.RoomID,
			Username: cl.Username,
		}

		hub.Broadcast <- msg
	}
}
