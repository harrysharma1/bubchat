package server

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHubRegisterUnregister(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	hub := NewHub()
	go hub.Run(ctx)

	client := &Client{
		send: make(chan Message, 1),
	}

	hub.register <- client
	time.Sleep(100 * time.Millisecond)
	assert.True(t, hub.clients[client])

	hub.unregister <- client
	time.Sleep(100 * time.Millisecond)
	assert.False(t, hub.clients[client])

}

func TestHubBroadcast(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	hub := NewHub()
	go hub.Run(ctx)

	c1 := &Client{
		send: make(chan Message, 1),
	}

	c2 := &Client{
		send: make(chan Message, 1),
	}

	hub.register <- c1
	hub.register <- c2

	<-c1.send
	<-c2.send

	time.Sleep(100 * time.Millisecond)

	expectedMessage := Message{
		Type:  "welcome",
		Value: "joined chat",
	}

	hub.broadcast <- expectedMessage
	msg1 := <-c1.send
	msg2 := <-c2.send

	assert.Equal(t, expectedMessage.Value, msg1.Value)
	assert.Equal(t, expectedMessage.Value, msg2.Value)

}
