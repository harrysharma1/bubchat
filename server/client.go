package server

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Constants to alter buffer/wait time for server read/write.
const (
	bufferSize = 2048
	wait       = 60 * time.Second

	pongWait  = wait
	pingWait  = (wait * 9) / 10
	writeWait = wait - (50 * time.Second)
)

var upgrader = websocket.Upgrader{
	WriteBufferSize: bufferSize,
	ReadBufferSize:  bufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // Upgrading standard HTTP to Websocket

// Message that will be passed to all clients via the Hub pattern.
type Message struct {
	Type       string    `json:"type"`        // Message type to decide how clientside code will handle the message.
	Value      string    `json:"value"`       // Value of the message e.g. message body.
	UserId     string    `json:"user_id"`     // UUID that helps id client's
	Username   string    `json:"username"`    // Username for clientside code.
	UploadTime time.Time `json:"upload_time"` // Upload time of message.
}

// Client to be used within the Hub logic of code.
type Client struct {
	hub      *Hub            // Hub for managing the clients connections.
	conn     *websocket.Conn // Client's websocket connection.
	send     chan Message    // Goroutine channel for concurrently sending message
	Username string          // Client username.
	UserId   string          // Client userid.
}

// Writing message to all clients through websocket hub.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingWait)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {

		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			err := c.conn.WriteJSON(message)
			if err != nil {
				log.Println(err)
				c.hub.unregister <- c
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}

	}

}

// Making sure all othe clients recieve message from current client.
func (c *Client) ReadPump() {
	defer func() {
		c.hub.broadcast <- Message{
			Type:       "exit",
			Value:      "left chat",
			UserId:     c.UserId,
			Username:   c.Username,
			UploadTime: time.Now(),
		}
		c.hub.unregister <- c
	}()
	c.conn.SetReadLimit(bufferSize / 2)
	c.conn.SetReadDeadline(time.Now().Add(wait))
	c.conn.SetPongHandler(func(appData string) error { c.conn.SetReadDeadline(time.Now().Add(wait)); return nil })

	for {
		var message Message
		err := c.conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println(err)
			}
			return
		}

		message.UserId = c.UserId
		message.Username = c.Username
		message.UploadTime = time.Now()
		c.hub.broadcast <- message
	}
}

func ServeWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		username = "dokja"
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		hub:      hub,
		conn:     conn,
		send:     make(chan Message, bufferSize/2),
		Username: username,
		UserId:   uuid.NewString(),
	}
	client.hub.register <- client

	go client.WritePump()
	go client.ReadPump()
}
