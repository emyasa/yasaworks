// Package tui implements the terminal user interface for Yasaworks. It's considered the lowest level module in tui package orchastrating other higher level modules.
package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/emyasa/yasaworks/internal/tui/blog"
	"github.com/emyasa/yasaworks/internal/tui/chat"
	"github.com/emyasa/yasaworks/internal/tui/splash"
	"github.com/emyasa/yasaworks/internal/tui/terms"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type page = int
const (
	splashPage page = iota
	blogPage
	termsPage
	chatPage
)

type model struct {
	theme theme.Theme
	page page
	splash splash.Model
	blog blog.Model
	terms terms.Model
	chat chat.Model

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
		page: splashPage,
		splash: splash.NewModel(basicTheme),
		blog: blog.NewModel(basicTheme, widthContainer, heightContainer),
		terms: terms.NewModel(basicTheme, widthContainer, heightContainer),
		chat: chat.NewModel(basicTheme),
		widthContainer: widthContainer,
		heightContainer: heightContainer,
	}, nil
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.splash.Init())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewportWidth = msg.Width
		m.viewportHeight = msg.Height
	case splash.SplashCompleteMsg:
		m.page = blogPage
	}

	cmd := m.hotKeyUpdate(msg)
	if cmd != nil {
		return m, cmd
	}

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
	case chatPage:
		var cmd tea.Cmd
		m.chat, cmd = m.chat.Update(msg)
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

