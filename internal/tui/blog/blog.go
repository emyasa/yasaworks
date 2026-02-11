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
	menuWidth int
	selected int
}

type blogEntry struct {
	name string
	viewport *viewport.Model
}

//go:embed entries/*.md
var entriesFS embed.FS
var blogEntries = []blogEntry{
	{name: "First Entry"},
}

func (m Model) Init() tea.Cmd {
	m.menuWidth = maxEntryWidth(blogEntries)
	contentWidth := 80 - m.menuWidth

	r, _ := glamour.NewTermRenderer(
	glamour.WithAutoStyle(),
	glamour.WithWordWrap(contentWidth))

	firstEntryContent, err := entriesFS.ReadFile("entries/first.md")
	if err != nil {
		panic(err)
	}

	detailContent, _ := r.Render(string(firstEntryContent))
	vp := viewport.New(contentWidth, 10)
	vp.SetContent(detailContent)

	blogEntries[0].viewport = &vp
	
	return nil
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
	nav := m.navView(entries, selected)
	containerWidth := 80 - m.menuWidth

	navStyle := m.Theme.Base().
		Width(45).
		Align(lipgloss.Right)

	vp := entries[selected].viewport
	content := lipgloss.JoinVertical(
		lipgloss.Top,
		vp.View(),
		navStyle.Render(nav),
	)

	containerStyle := m.Theme.Base().
		Width(containerWidth).
		MarginTop(1).
		Padding(0, 1)

	return containerStyle.Render(content)
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

