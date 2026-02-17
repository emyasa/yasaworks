// Package cursor
package cursor

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type Model struct {
	Theme theme.Theme
	visible bool
}

type CursorTickMsg struct {}

func (m Model) Init() tea.Cmd {
	return tea.Every(time.Millisecond*700, func(t time.Time) tea.Msg {
		return CursorTickMsg{}
	})
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg.(type) {
	case CursorTickMsg:
		m.visible = !m.visible
		return m, tea.Every(time.Millisecond*700, func(t time.Time) tea.Msg {
			return CursorTickMsg{}
		})
	}

	return m, nil
}

func (m Model) View() string {
	if m.visible {
		return m.Theme.Base().
			Background(m.Theme.Brand()).
			Render(" ")
	}

	return " "
}

