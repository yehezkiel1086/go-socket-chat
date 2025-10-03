package domain

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
