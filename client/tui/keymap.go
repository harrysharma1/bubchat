package tui

import "github.com/charmbracelet/bubbles/key"

// map of special keys in TUI.
type KeyMap struct {
	Send key.Binding // key for sending text.
	Quit key.Binding // key for quitting as client.
}

/*
Interface implementation of the key binding ShortHelp() function.

(note: will fail if thes interfaces aren't implemented)
*/
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Send}
}

/*
Interface implementation of the key binding FullHelp() function.

(note: will fail if thes interfaces aren't implemented)
*/
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit, k.Send},
	}
}

var keys = KeyMap{
	Send: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "send message"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
} // Keys to be passed into ChatModel.
