// Package terms handles all the man interface page
package terms

import (
	_ "embed"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type Model struct {
	theme theme.Theme
	viewport *viewport.Model
}

//go:embed man.md
var manContent string

func NewModel(theme theme.Theme, contentWidth int, contentHeight int) Model {
	r, _ := glamour.NewTermRenderer(
		theme.MarkdownStyle(),
		glamour.WithPreservedNewLines())

	content, _ := r.Render(manContent)

	vp := viewport.New(contentWidth, contentHeight - 9)
	vp.SetContent(content)

	return Model{
		theme: theme,
		viewport: &vp,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "j", "down":
			m.viewport.ScrollDown(1)
		case "shift+tab", "k", "up":
			m.viewport.ScrollUp(1)
		}
	}

	return m, nil
}

func (m Model) View() string {
	termsContent := m.theme.Base().
		Margin(1, 1, 0).
		Render(m.viewport.View())

	cursor := m.theme.Base().
		Margin(0, 1).
		Render(":")

	return lipgloss.JoinVertical(
		lipgloss.Top,
		termsContent,
		cursor,
	)
}

