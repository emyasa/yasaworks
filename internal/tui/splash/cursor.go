package splash

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type cursorState struct {
	visible bool
}

type cursorTickMsg struct {}

func (m Model) cursorInit() tea.Cmd {
	return tea.Every(time.Millisecond*700, func(t time.Time) tea.Msg {
		return cursorTickMsg{}
	})
}

func (m Model) cursorUpdate(msg tea.Msg) (Model, tea.Cmd) {
	switch msg.(type) {
	case cursorTickMsg:
		m.state.cursor.visible = !m.state.cursor.visible
		return m, tea.Every(time.Millisecond*700, func(t time.Time) tea.Msg {
			return cursorTickMsg{}
		})
	}

	return m, nil
}

func (m Model) cursorView() string {
	if m.state.cursor.visible {
		return m.Theme.Base().
			Background(m.Theme.Brand()).
			Render(" ")
	}

	return " "
}

