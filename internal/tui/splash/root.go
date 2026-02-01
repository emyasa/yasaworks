// Package splash handles the tui's splash state
package splash

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	state state
	viewportWidth int
	viewportHeight int
}

type state struct {
	cursor cursorState
}

func(m Model) Init() tea.Cmd {
	return m.cursorInit()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {		
	case tea.WindowSizeMsg:
		m.viewportWidth = msg.Width
		m.viewportHeight = msg.Height
	case cursorTickMsg:
		m, cmd := m.cursorUpdate(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	return lipgloss.Place(
		m.viewportWidth,
		m.viewportHeight,
		lipgloss.Center,
		lipgloss.Center,
		m.logoView(),
	)
}

