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

func(m Model) SplashInit() tea.Cmd {
	return m.cursorInit()
}

func (m Model) SplashUpdate(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {		
	case tea.WindowSizeMsg:
		m.viewportWidth = msg.Width
		m.viewportHeight = msg.Height
	case CursorTickMsg:
		m, cmd := m.CursorUpdate(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) SplashView() string {
	return lipgloss.Place(
		m.viewportWidth,
		m.viewportHeight,
		lipgloss.Center,
		lipgloss.Center,
		m.logoView(),
	)
}

