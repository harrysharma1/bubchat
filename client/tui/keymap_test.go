package tui

import (
	"testing"

	"github.com/charmbracelet/bubbles/key"
	"github.com/stretchr/testify/assert"
)

func TestShortHelp(t *testing.T) {
	expected := []key.Binding{keys.Quit, keys.Send}
	actual := keys.ShortHelp()
	assert.Equal(t, expected, actual)
}

func TestFullHelp(t *testing.T) {
	expected := [][]key.Binding{
		{keys.Quit, keys.Send},
	}
	actual := keys.FullHelp()
	assert.Equal(t, expected, actual)
}
