package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootPersistentVersionFlag(t *testing.T) {
	resetGlobalState(t)
	cmd := rootCmd
	cmd.SetArgs(
		[]string{
			"--version",
		},
	)

	out := captureStdout(t, func() {
		err := cmd.Execute()
		assert.Nil(t, err)
	})
	assert.Equal(t, EXPECTED_BUB_CHAT_VERSION, out)
}
