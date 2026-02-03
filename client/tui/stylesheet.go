package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	timeStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#BC13FE"))
	userStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#18d925"))
	systemStyle = lipgloss.NewStyle().Italic(true).Faint(true)
)
