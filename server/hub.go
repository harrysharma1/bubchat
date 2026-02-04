package server

// Hub for managing clients.
type Hub struct {
	clients    map[*Client]bool // Goroutine channel for keeping track of connected clients.
	broadcast  chan Message     // Goroutine channel for broadcasting message to all other clients.
	register   chan *Client     // Goroutine channel for registering client to hub.
	unregister chan *Client     // Gorouting channel for unregistering client from hub.
}

// Initialising a new hub.
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

/*
Programm loop for hub. Either you:

- register client (and send welcome message to all other clients).

- unregister client.

- broadcast messages to all clients.
*/
func (h *Hub) Run() {
	for {

		select {
		case client := <-h.register:
			h.clients[client] = true

			message := Message{
				Type:     "welcome",
				Value:    "joined chat",
				Username: client.Username,
				UserId:   client.UserId,
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
