// Package tui implements the terminal user interface for Yasaworks. It's considered the lowest level module in tui package orchastrating other higher level modules.
package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/emyasa/yasaworks/internal/tui/blog"
	"github.com/emyasa/yasaworks/internal/tui/splash"
	"github.com/emyasa/yasaworks/internal/tui/terms"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type page = int
const (
	splashPage page = iota
	blogPage
	termsPage
)

type model struct {
	theme theme.Theme
	splash splash.Model
	blog blog.Model
	terms terms.Model

	page page
	viewportWidth int
	viewportHeight int
	widthContainer int
	heightContainer int
}

func NewModel() (tea.Model, error) {
	basicTheme := theme.BasicTheme()
	widthContainer := 80
	heightContainer := 30

	return model{
		theme: basicTheme,
		splash: splash.Model{ Theme: basicTheme },
		blog: blog.NewModel(basicTheme, widthContainer, heightContainer),
		terms: terms.NewModel(basicTheme, widthContainer, heightContainer),
		page: splashPage,
		widthContainer: widthContainer,
		heightContainer: heightContainer,
	}, nil
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.splash.Init())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.viewportWidth = msg.Width
		m.viewportHeight = msg.Height
	case splash.SplashCompleteMsg:
		m.page = blogPage
	}

	m.headerUpdate(msg)
	switch m.page {
	case splashPage:
		var cmd tea.Cmd
		m.splash, cmd = m.splash.Update(msg)
		return m, cmd
	case blogPage:
		var cmd tea.Cmd
		m.blog, cmd = m.blog.Update(msg)
		return m, cmd
	case termsPage:
		var cmd tea.Cmd
		m.terms, cmd = m.terms.Update(msg)
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
			m.theme.Base().
				MaxWidth(m.widthContainer).
				MaxHeight(m.heightContainer).
				Render(child),
		)
}

