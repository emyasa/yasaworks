// Package chat handles conversation between potential client and the company
package chat

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type Model struct {
	theme theme.Theme
	input textinput.Model
}

func NewModel(theme theme.Theme) Model {
	ti := textinput.New()
	ti.Prompt = "> "
	ti.Placeholder = "type a command..."
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 60

	// Remove default styling for clean terminal look
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))


	return Model{
		theme: theme,
		input: ti,
	}
}

func (m Model) View() string {
	return "FOCUSED: " + fmt.Sprint(m.input.Focused()) + "\n\n" + m.input.View()
}

