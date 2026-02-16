package server

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestServeWSRegisterClient(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	hub := NewHub()
	go hub.Run(ctx)

	server := startTestServer(hub)
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws?username=testuser"

	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	assert.NoError(t, err)
	defer ws.Close()

	var msg Message
	err = ws.ReadJSON(&msg)

	assert.NoError(t, err)
	assert.Equal(t, "welcome", msg.Type)
	assert.Equal(t, "joined chat", msg.Value)
	assert.Equal(t, "testuser", msg.Username)
}

func TestWritePumpSendMessage(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	hub := NewHub()
	go hub.Run(ctx)

	server := startTestServer(hub)
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws?username=testuser"

	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	assert.NoError(t, err)
	defer ws.Close()

	ws.SetReadDeadline(time.Now().Add(2 * time.Second))

	var welcome Message
	err = ws.ReadJSON(&welcome)
	assert.NoError(t, err)
	assert.Equal(t, "welcome", welcome.Type)

	err = ws.WriteJSON(Message{
		Type:  "chat",
		Value: "Hello, World!",
	})
	assert.NoError(t, err)

	var received Message
	err = ws.ReadJSON(&received)
	assert.NoError(t, err)

	assert.Equal(t, "chat", received.Type)
	assert.Equal(t, "Hello, World!", received.Value)
	assert.Equal(t, "testuser", received.Username)
}

func TestWritePumpTextMessageWebsocketType(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	hub := NewHub()
	go hub.Run(ctx)

	server := startTestServer(hub)
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws?username=testuser"

	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	assert.NoError(t, err)
	defer ws.Close()

	ws.SetReadDeadline(time.Now().Add(pingWait + time.Second))
	msgType, _, err := ws.ReadMessage()
	assert.NoError(t, err)
	assert.Equal(t, websocket.TextMessage, msgType)
}
