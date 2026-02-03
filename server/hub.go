package server

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {

		select {
		case client := <-h.register:
			h.clients[client] = true

			message := Message{
				Type:     "welcome",
				Value:    "joined chat",
				Username: client.username,
				UserId:   client.userId,
			}

			for client := range h.clients {
				client.send <- message
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		case message := <-h.broadcast:
			for client := range h.clients {
				client.send <- message
			}
		}

	}
}
