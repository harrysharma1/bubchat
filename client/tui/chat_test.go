package tui

import (
	"bubchat/client/ws"
	"io"
	"testing"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/exp/teatest"
)

func MockInitialChatModel() *ChatModel {
	cTextArea := textarea.New()
	cTextArea.Placeholder = "Send a message..."
	cTextArea.Focus()

	cTextArea.Prompt = "┃ "
	cTextArea.CharLimit = 320

	cTextArea.SetWidth(30)
	cTextArea.SetHeight(3)

	cTextArea.FocusedStyle.CursorLine = lipgloss.NewStyle()
	cTextArea.ShowLineNumbers = false

	mViewPort := viewport.New(35, 5)

	cTextArea.KeyMap.InsertNewline.SetEnabled(false)

	return &ChatModel{
		viewPort:     mViewPort,
		messages:     []string{},
		chatTextArea: cTextArea,
		err:          nil,
		help:         help.New(),
		keys:         keys,
		Client: &ws.Client{
			Username: "john",
			UserId:   "abcdefg",
		},
	}

}

func TestTUIFullOutput(t *testing.T) {
	m := MockInitialChatModel()
	m.Client = &ws.Client{
		Username: "test",
		UserId:   "abcdef",
	}
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(300, 100))
	tm.Send(tea.KeyMsg{Type: tea.KeyCtrlC})
	out, err := io.ReadAll(tm.FinalOutput(t))
	if err != nil {
		t.Error(err)
	}
	teatest.RequireEqualOutput(t, out)
}
