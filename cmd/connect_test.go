package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultClientPort(t *testing.T) {
	resetGlobalState(t)
	cmd := rootCmd
	cmd.SetArgs(
		[]string{
			"connect",
		},
	)

	cmd.Execute()
	expectedClientPort := "6969"
	assert.Equal(t, expectedClientPort, clientPort)
}

func TestDefaultClientHost(t *testing.T) {
	resetGlobalState(t)
	cmd := rootCmd
	cmd.SetArgs(
		[]string{
			"connect",
		},
	)

	cmd.Execute()
	expectedClientHost := "127.0.0.1"
	assert.Equal(t, expectedClientHost, clientHost)

}

func TestDefaultClientName(t *testing.T) {
	resetGlobalState(t)
	cmd := rootCmd
	cmd.SetArgs(
		[]string{
			"connect",
		},
	)

	cmd.Execute()
	expectedClientName := "dokja"
	assert.Equal(t, expectedClientName, clientName)

}

func TestClientFailsGracefullyIfServerIsDown(t *testing.T) {
	resetGlobalState(t)
	cmd := rootCmd
	cmd.SetArgs(
		[]string{
			"connect",
		},
	)

	err := cmd.Execute()
	assert.ErrorContains(t, err, "dial tcp 127.0.0.1:6969: connect: connection refused")
}

func TestConnectPersistentVersionFlag(t *testing.T) {
	resetGlobalState(t)
	cmd := rootCmd
	cmd.SetArgs(
		[]string{
			"connect",
			"--version",
		},
	)

	out := captureStdout(t, func() {
		err := cmd.Execute()
		assert.Nil(t, err)
	})
	assert.Equal(t, EXPECTED_BUB_CHAT_VERSION, out)
}
