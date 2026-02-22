// Package chat handles conversation between potential client and the company
package chat

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type mode = int
const (
	Normal mode = iota
	Insert
)

type Model struct {
	theme theme.Theme
	input textinput.Model
	Mode mode
}

func NewModel(theme theme.Theme) Model {
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
	}
}

func (m *Model) Init() tea.Cmd {
	m.Mode = Insert
	m.input.Focus()
	return m.input.Cursor.BlinkCmd()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
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
	inputView := m.theme.Base().
		MarginLeft(1).
		Render(m.input.View())

	modeString := ""
	if m.Mode == Insert {
		modeString = "-- INSERT --"
	}

	statusLine := m.theme.Base().
		MarginLeft(1).
		Render(modeString)

	child := lipgloss.JoinVertical(lipgloss.Left, inputView, statusLine)

	return lipgloss.Place(80, 24, lipgloss.Left, lipgloss.Bottom, child)
}

