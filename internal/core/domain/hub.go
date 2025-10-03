package domain

import "fmt"

type Hub struct {
	Broadcast  chan *Message    `json:"broadcast"`
	Register   chan *Client     `json:"register"`
	Unregister chan *Client     `json:"unregister"`
	Rooms      map[string]*Room `json:"rooms"`
}

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					return
				}

				h.Rooms[cl.RoomID].Clients[cl.ID] = cl
			}
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						msg := &Message{
							Content: fmt.Sprintf("%v has left the chat room", cl.Username),
							RoomID: cl.RoomID,
							Username: cl.Username,
						}

						h.Broadcast <- msg
					}

					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)
				}
			}
		case msg := <-h.Broadcast:
			if _, ok := h.Rooms[msg.RoomID]; ok {
				for _, cl := range h.Rooms[msg.RoomID].Clients {
					cl.Message <- msg
				}
			}
		}
	}
}
