// Package blog handles all the blog-related entries
package blog

import (
	"embed"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type Model struct {
	Theme theme.Theme
	ContainerWidth int

	menuWidth int
	contentWidth int
	navWidth int
	selected int
}

type blogEntry struct {
	name string
	mdPath string
	viewport *viewport.Model
}

//go:embed entries/*.md
var entriesFS embed.FS
var blogEntries = []blogEntry{
	{name: "First Entry", mdPath: "entries/first.md"},
}

func NewModel(theme theme.Theme, containerWidth int) Model {
	menuWidth := maxEntryWidth(blogEntries)
	contentWidth := containerWidth - menuWidth
	navWidth := contentWidth - 8 

	r, _ := glamour.NewTermRenderer(
	glamour.WithAutoStyle(),
	glamour.WithWordWrap(contentWidth))

	for _, entry := range blogEntries {
		content, err := entriesFS.ReadFile(entry.mdPath)
		if err != nil {
			panic(err)
		}

		detailContent, _ := r.Render(string(content))
		vp := viewport.New(contentWidth, 10)
		vp.SetContent(detailContent)

		blogEntries[0].viewport = &vp
	}
	
	return Model{
		Theme: theme,
		menuWidth: menuWidth,
		contentWidth: contentWidth,
		navWidth: navWidth,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "j", "down":
			if m.selected < len(blogEntries) - 1 {
				m.selected += 1
			}

			return m, nil
		case "shift+tab", "k", "up":
			if m.selected > 0 {
				m.selected -= 1
			}
			
			return m, nil
		case "n":
			vp := blogEntries[m.selected].viewport
			vp.ScrollDown(10)

			return m, nil
		case "N":
			vp := blogEntries[m.selected].viewport
			vp.ScrollUp(10)

			return m, nil
		}
	}

	return m, nil
}

func (m Model) View() string {
	menuContent := m.renderBlogMenu(blogEntries, m.selected)
	detailContent := m.renderBlogDetail(blogEntries, m.selected)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		menuContent,
		"  ",
		detailContent,
	)
}

func (m Model) renderBlogMenu(entries []blogEntry, selected int) string {
	m.menuWidth = maxEntryWidth(entries)

	var sb strings.Builder
	for i, e := range entries {
		menuItemStyle := m.Theme.Base().
			Width(m.menuWidth + 2).
			Padding(0, 1)

		if i == selected {
			menuItemStyle = menuItemStyle.Background(m.Theme.Highlight()).
				Foreground(m.Theme.Accent()).
				Bold(true)
		}

		sb.WriteString(menuItemStyle.Render(e.name))
		if i < len(entries) - 1 {
			sb.WriteString("\n")
		}
	}

	containerStyle := m.Theme.Base().
		MarginTop(1).
		Padding(0, 1)

	return containerStyle.Render(sb.String())
}

func (m Model) renderBlogDetail(entries []blogEntry, selected int) string {
	vp := entries[selected].viewport
	content := lipgloss.JoinVertical(
		lipgloss.Top,
		vp.View(),
		m.navView(entries, selected),
	)

	return m.Theme.Base().Render(content)
}

func maxEntryWidth(entries []blogEntry) int {
	max := 0
	for _, e := range entries {
		if w := lipgloss.Width(e.name); w > max {
			max = w
		}
	}

	return max
}

