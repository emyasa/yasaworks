// Package tui implements the terminal user interface for Yasaworks
package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case tea.WindowSizeMsg:
		m.viewportWidth = msg.Width
		m.viewportHeight = msg.Height
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
	default:
		header := m.HeaderView()
		footer := m.FooterView()

		items := []string{}
		items = append(items, header)
		items = append(items, footer)

		child := lipgloss.JoinVertical(
			lipgloss.Center,
			items...,
		)

		return lipgloss.Place(
			m.viewportWidth,
			m.viewportHeight,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.NewStyle().Render(child),
		)
	}
}

