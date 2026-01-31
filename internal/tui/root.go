// Package tui implements the terminal user interface for Yasaworks
package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	state state
	viewportWidth int
	viewportHeight int
}

type state struct {
	cursor cursorState
}

func NewModel() (tea.Model, error) {
	return model{}, nil
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.CursorInit())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case tea.WindowSizeMsg:
		m.viewportWidth = msg.Width
		m.viewportHeight = msg.Height
	case CursorTickMsg:
		m, cmd := m.CursorUpdate(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	return lipgloss.Place(
		m.viewportWidth,
		m.viewportHeight,
		lipgloss.Center,
		lipgloss.Center,
		m.LogoView(),
	)
}

