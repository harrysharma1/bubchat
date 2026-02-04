package ws

import (
	"bubchat/server"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

// Constants to alter buffer/wait time for client read.
const (
	bufferSize = 2048
	wait       = 60 * time.Second

	pongWait = wait
	pingWait = (wait * 9) / 10
)

// User facing client.
type Client struct {
	Conn     *websocket.Conn // Websocket connection for client to access websocket server.
	Username string          // Client's username for tui use case.
	UserId   string          // Client's userid for websocket connection.
	Program  *tea.Program    // Program allows the to initialise the tea program state.
}

// Create a new client with a websocket connection.
func NewClient(url, username string, p *tea.Program) (*Client, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	return &Client{
		Conn:     conn,
		Username: username,
		Program:  p,
	}, nil
}

// Reading message from websocket server to be sent to TUI processing.
func (c *Client) ReadPump() {
	defer c.Conn.Close()

	c.Conn.SetReadDeadline(time.Now().Add(wait))
	c.Conn.SetPingHandler(func(appData string) error { c.Conn.SetReadDeadline(time.Now().Add(wait)); return nil })

	for {
		var message server.Message

		err := c.Conn.ReadJSON(&message)
		if err != nil {
			return
		}

		if message.Type == "welcome" {
			c.UserId = message.UserId
			c.Program.Send(message)
			continue
		}
		c.Program.Send(message)
	}
}

func (c *Client) Run() {
	go c.ReadPump()
}
