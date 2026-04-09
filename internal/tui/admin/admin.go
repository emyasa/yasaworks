// Package admin
package admin

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/emyasa/yasaworks/internal/registry"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type mode = int
const (
	Normal mode = iota
	Insert
)

type Model struct {
	Mode mode
	theme theme.Theme
	input textinput.Model
	conversations []conversation
	conversationsIndex map[string]int
	selectedConversationIndex int
	messages []string
	conn *registry.Connection
}

func NewModel(theme theme.Theme, conn *registry.Connection) Model {
	ti := textinput.New()
	ti.Prompt = "> "
	ti.Placeholder = "type a message..."

	// Remove default styling for clean terminal look
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))

	return Model{
		theme: theme,
		input: ti,
		conn: conn,
		conversationsIndex: map[string]int{},
	}
}

type clientMessageEvent registry.MessageEvent

func readFromChannel(conn *registry.Connection) tea.Cmd {
	if conn == nil {
		return nil
	}

	return func() tea.Msg {
		return clientMessageEvent(conn.FetchMessage())
	}
}

func (m *Model) Init() tea.Cmd {
	m.Mode = Insert
	m.input.Focus()
	return tea.Batch(m.input.Cursor.BlinkCmd(), readFromChannel(m.conn))
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case clientMessageEvent:
		cMsgEvent := registry.MessageEvent(msg)
		m.updateConversations(cMsgEvent)
		return m, readFromChannel(m.conn)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			if m.Mode == Insert {
				m.Mode = Normal
				m.input.Cursor.SetMode(cursor.CursorStatic)
			}
		case "i":
			if m.Mode == Normal {
				m.Mode = Insert
				m.input.Cursor.SetMode(cursor.CursorBlink)

				return m, m.input.Cursor.BlinkCmd()
			}
		}
	}

	if m.Mode == Normal {
		return m, nil
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.conversationsView(),
		m.chatPanelView(),
	)
}

