// Package tui implements the terminal user interface for Yasaworks. It's considered the lowest level module in tui package orchastrating other higher level modules.
package tui

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/emyasa/yasaworks/internal/db"
	"github.com/emyasa/yasaworks/internal/registry"
	"github.com/emyasa/yasaworks/internal/tui/admin"
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
	adminPage
)

type model struct {
	db *db.DB
	fingerprint string
	anonymous bool
	isAdmin bool
	clientIP *string

	theme theme.Theme
	page page
	splash splash.Model
	blog blog.Model
	terms terms.Model
	chat chat.Model
	admin *admin.Model

	viewportWidth int
	viewportHeight int
	widthContainer int
	heightContainer int
}

func NewModel(
	ctx context.Context,
	db *db.DB,
	fingerprint string,
	anonymous bool,
	isAdmin bool,
	clientIP *string,
	conn *registry.Connection,
) (tea.Model, error) {
	basicTheme := theme.BasicTheme()
	widthContainer := 80
	heightContainer := 30

	return model{
		db: db,
		fingerprint: fingerprint,
		anonymous: anonymous,
		isAdmin: isAdmin,
		clientIP: clientIP,
		theme: basicTheme,
		page: splashPage,
		splash: splash.NewModel(basicTheme),
		blog: blog.NewModel(basicTheme, widthContainer, heightContainer),
		terms: terms.NewModel(basicTheme, widthContainer, heightContainer),
		chat: chat.NewModel(ctx, db, basicTheme, conn),
		admin: admin.NewModel(basicTheme, conn),
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
		m.admin.ViewportWidth = msg.Width
		m.admin.ViewportHeight = msg.Height
	case splash.SplashCompleteMsg:
		if m.isAdmin {
			m.page = adminPage
			return m, m.admin.Init()
		}

		m.page = blogPage
	}

	cmd, handled := m.hotKeyUpdate(msg)
	if handled {
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
	case adminPage:
		var cmd tea.Cmd
		m.admin, cmd = m.admin.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	switch m.page {
	case splashPage:
		return m.splash.View()
	case adminPage:
		return m.admin.View()
	default:
		return m.layout(
			m.headerView(),
			m.getContent(),
			m.footerView(),
		)
	}
}

