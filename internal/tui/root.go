// Package tui implements the terminal user interface for Yasaworks
package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/emyasa/yasaworks/internal/tui/splash"
)

type page = int
const (
	splashPage page = iota
	blogPage
)

type model struct {
	page page
	splash splash.Model
	viewportWidth int
	viewportHeight int
}

func NewModel() (tea.Model, error) {
	return model{
		page: splashPage,
	}, nil
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.splash.Init())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case splash.SplashCompleteMsg:
		m.page = blogPage
	}

	if m.page == splashPage {
		var cmd tea.Cmd
		m.splash, cmd = m.splash.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	switch m.page {
	case splashPage:
		return m.splash.View()
	}

	return ""
}

