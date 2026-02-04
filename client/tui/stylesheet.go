package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	timeStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#BC13FE")) // Styling for time.
	userStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#18d925")) // Styling for user (colour changes based on uuid in Update()).
	systemStyle  = lipgloss.NewStyle().Italic(true).Faint(true)              // Styling for message as system.
	messageStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")) // Styling for message as client.
)
