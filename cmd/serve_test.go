package cmd

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultServerPort(t *testing.T) {
	resetGlobalState(t)

	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan struct{})
	go func() {
		defer close(done)
		err := runServer(ctx)
		assert.Nil(t, err)
	}()
	time.Sleep(2 * time.Second)
	cancel()
	<-done
	expectedServerPort := "6969"
	assert.Equal(t, expectedServerPort, serverPort)
}

func TestDefaultServerHost(t *testing.T) {
	resetGlobalState(t)

	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan struct{})
	go func() {
		defer close(done)
		err := runServer(ctx)
		assert.Nil(t, err)

	}()
	time.Sleep(2 * time.Second)
	cancel()
	<-done
	expectedServerHost := "127.0.0.1"
	assert.Equal(t, expectedServerHost, serverHost)

}

func TestServerLogs(t *testing.T) {
	resetGlobalState(t)

	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan struct{})
	var out string
	go func() {
		defer close(done)
		out = captureLog(t, func() {
			err := runServer(ctx)
			assert.Nil(t, err)
		})
		fmt.Println(out)
	}()
	time.Sleep(2 * time.Second)
	cancel()
	<-done
	assert.Contains(t, out, "INFO Websocket server launched on ws::/127.0.0.1:6969/ws")
	assert.Contains(t, out, "INFO Shutting down server gracefully...")
	assert.Contains(t, out, "INFO Server exited cleanly 👋")
}

func TestServePersistentVersionFlag(t *testing.T) {
	resetGlobalState(t)
	cmd := rootCmd
	cmd.SetArgs(
		[]string{
			"serve",
			"--version",
		},
	)

	out := captureStdout(t, func() {
		err := cmd.Execute()
		assert.Nil(t, err)
	})
	assert.Equal(t, EXPECTED_BUB_CHAT_VERSION, out)
}
