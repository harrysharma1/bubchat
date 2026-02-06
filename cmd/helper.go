package cmd

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/charmbracelet/log"
)

const EXPECTED_BUB_CHAT_VERSION = "bubchat (v0.0.1)"

func resetGlobalState(t *testing.T) {
	t.Helper()
	clientHost = ""
	clientName = ""
	clientPort = ""
	serverHost = ""
	serverPort = ""
}

func captureLog(t *testing.T, fn func()) string {
	t.Helper()

	var buffer bytes.Buffer

	org := log.Default()
	logger := log.New(&buffer)
	log.SetDefault(logger)

	fn()

	log.SetDefault(org)
	return buffer.String()
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}
