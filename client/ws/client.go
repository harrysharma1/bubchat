package ws

import (
	"bubchat/server"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

const (
	bufferSize = 2048
	wait       = 60 * time.Second

	pongWait  = wait
	pingWait  = (wait * 9) / 10
	writeWait = wait - (50 * time.Second)
)

type Client struct {
	conn     *websocket.Conn
	username string
	userId   string
	program  *tea.Program
}

func NewClient(url, username string, p *tea.Program) (*Client, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:     conn,
		username: username,
		program:  p,
	}, nil
}

func (c *Client) ReadPump() {
	defer c.conn.Close()

	c.conn.SetReadDeadline(time.Now().Add(wait))
	c.conn.SetPingHandler(func(appData string) error { c.conn.SetReadDeadline(time.Now().Add(wait)); return nil })

	for {
		var message server.Message

		err := c.conn.ReadJSON(&message)
		if err != nil {
			return
		}

		if message.Type == "welcome" {
			c.userId = message.UserId
			continue
		}

	}
}

func (c *Client) Run() {
	go c.ReadPump()
}
