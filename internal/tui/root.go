// Package tui implements the terminal user interface for Yasaworks. It's considered the lowest level module in tui package orchastrating other higher level modules.
package tui

import (
	"math"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/emyasa/yasaworks/internal/tui/blog"
	_ "github.com/emyasa/yasaworks/internal/tui/blog/entries"
	"github.com/emyasa/yasaworks/internal/tui/splash"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type page = int
const (
	splashPage page = iota
	blogPage
)

type model struct {
	splash splash.Model
	blog blog.Model

	page page
	viewportWidth int
	viewportHeight int
	widthContainer int
	heightContainer int
}

func NewModel() (tea.Model, error) {
	basicTheme := theme.BasicTheme()

	return model{
		splash: splash.Model{ Theme: basicTheme },
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
		m.widthContainer = 80
		m.heightContainer = int(math.Min(float64(msg.Height), 30))
	case splash.SplashCompleteMsg:
		m.page = blogPage
	}

	switch m.page {
	case splashPage:
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
		return m.layout(
			m.headerView(),
			m.getContent(),
			m.footerView(),
		)
	}
}

func (m model) layout(header, content, footer string) string {
		height := m.heightContainer
		height -= lipgloss.Height(header)
		height -= lipgloss.Height(footer)

		body := lipgloss.NewStyle().
			Width(80).
			Height(height).
			Render(content)

		child := lipgloss.JoinVertical(
			lipgloss.Left,
			header,
			body,
			footer,
		)

		return lipgloss.Place(
			m.viewportWidth,
			m.viewportHeight,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.NewStyle().
				MaxWidth(m.widthContainer).
				MaxHeight(m.heightContainer).
				Render(child),
		)
}

