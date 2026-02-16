// Package terms handles all the man interface page
package terms

import (
	_ "embed"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type Model struct {
	Theme theme.Theme
}

//go:embed man.md
var manContent string

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg.(type) {
	case tea.WindowSizeMsg:
	}

	return m, nil
}

func (m Model) View() string {
	r, _ := glamour.NewTermRenderer(
		m.Theme.MarkdownStyle(),
		glamour.WithPreservedNewLines())

	content, _ := r.Render(manContent)

	return m.Theme.Base().
		Margin(1).
		Render(content)
}

