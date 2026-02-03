package tui

import (
	"bubchat/client/ws"
	"bubchat/helper"
	"bubchat/server"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const gap = "\n\n"

type (
	ErrorMsg       error
	SocketErrorMsg error
)

type ChatModel struct {
	viewPort     viewport.Model
	messages     []string
	chatTextArea textarea.Model
	Client       *ws.Client
	err          error
	help         help.Model
	keys         KeyMap
}

func InitialChatModel() *ChatModel {
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
	}

}

func (cm *ChatModel) Init() tea.Cmd {
	return textarea.Blink
}

func (cm *ChatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tAreaCmd tea.Cmd
		vPortCmd tea.Cmd
	)

	cm.chatTextArea, tAreaCmd = cm.chatTextArea.Update(msg)
	cm.viewPort, vPortCmd = cm.viewPort.Update(msg)

	switch msg := msg.(type) {

	case server.Message:
		incomingMessage := server.Message(msg)
		userStyle = userStyle.Foreground(lipgloss.Color(helper.HexFromUUID(helper.FirstN(msg.UserId, 6))))

		var line string
		time := timeStyle.Render(fmt.Sprintf("[%s]", incomingMessage.UploadTime.Format("15:04:05")))
		user := userStyle.Render(fmt.Sprintf("%s (%s)", incomingMessage.Username, helper.FirstN(incomingMessage.UserId, 6)))
		system := systemStyle.Render(incomingMessage.Value)

		switch incomingMessage.Type {
		case "welcome", "exit":
			line = fmt.Sprintf(
				"%s %s %s",
				time,
				user,
				system,
			)
		default:
			line = fmt.Sprintf(
				"%s %s: %s",
				time,
				user,
				system,
			)
		}
		cm.messages = append(cm.messages, line)
		cm.viewPort.SetContent(
			lipgloss.NewStyle().
				Width(cm.viewPort.Width).
				Render(strings.Join(cm.messages, "\n")),
		)
		cm.viewPort.GotoBottom()

	case tea.WindowSizeMsg:
		cm.viewPort.Width = msg.Width
		cm.chatTextArea.SetWidth(msg.Width)
		cm.viewPort.Height = msg.Height - cm.chatTextArea.Height() - lipgloss.Height(gap)

		if len(cm.messages) > 0 {
			// Wrap content before setting it.
			cm.viewPort.SetContent(lipgloss.NewStyle().Width(cm.viewPort.Width).Render(strings.Join(cm.messages, "\n")))
		}
		cm.viewPort.GotoBottom()

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			fmt.Println(cm.chatTextArea.Value())
			return cm, tea.Quit
		case tea.KeyEnter:
			text := strings.TrimSpace(cm.chatTextArea.Value())
			if text == "" {
				return cm, nil
			}

			out := server.Message{
				Value:      text,
				UserId:     cm.Client.UserId,
				Username:   cm.Client.Username,
				UploadTime: time.Now(),
			}
			if err := cm.Client.Conn.WriteJSON(out); err != nil {
				return cm, func() tea.Msg {
					return ErrorMsg(err)
				}
			}
			cm.chatTextArea.Reset()

		}
	case ErrorMsg:
		cm.err = msg
		return cm, nil
	}
	return cm, tea.Batch(tAreaCmd, vPortCmd)
}

func (cm *ChatModel) View() string {
	user := userStyle.Render(fmt.Sprintf("%s (%s)", cm.Client.Username, helper.FirstN(cm.Client.UserId, 6)))
	return fmt.Sprintf(
		"%s%s%s%s%s - %s",
		cm.viewPort.View(),
		gap,
		cm.chatTextArea.View(),
		gap,
		user,
		cm.help.View(cm.keys),
	)
}
