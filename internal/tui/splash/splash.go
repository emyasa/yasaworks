// Package splash handles the tui's splash state
package splash

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type Model struct {
	Theme theme.Theme
	state state
	viewportWidth int
	viewportHeight int
}

type state struct {
	cursor cursorState
}

type SplashCompleteMsg struct {}

func(m Model) Init() tea.Cmd {
	completeSplashCmd := tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		return SplashCompleteMsg{}
	})

	return tea.Batch(completeSplashCmd, m.cursorInit())
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

