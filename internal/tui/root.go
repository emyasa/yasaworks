// Package tui implements the terminal user interface for Yasaworks
package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/emyasa/yasaworks/internal/tui/splash"
)

type model struct {
	splash splash.Model
	viewportWidth int
	viewportHeight int
}

func NewModel() (tea.Model, error) {
	return model{}, nil
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.splash.SplashInit())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	default:
		var cmd tea.Cmd
		m.splash, cmd = m.splash.SplashUpdate(msg)
		return m, cmd
	}
}

func (m model) View() string {
	return m.splash.SplashView()
}

