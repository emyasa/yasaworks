package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) getBlogMenuContent() string {
	menuWidth := 0
	pages := strings.Builder{}

	for _, p := range m.blogEntries {
		w := lipgloss.Width(getBlogEntryName(p))
		if w > menuWidth {
			menuWidth = w
		}
	}

	menuItem := lipgloss.NewStyle().
		Width(menuWidth+2).
		Padding(0, 1)

	for _, p := range m.blogEntries {
		name := getBlogEntryName(p)
		content := menuItem.Render(name)
		pages.WriteString(content + "\n")
	}

	return lipgloss.NewStyle().
		MarginTop(1).
		Padding(0, 1).
		Render(pages.String())
}

func getBlogEntryName(entry blogEntry) string {
	switch entry {
	case firstEntry:
		return "First Entry"
	case secondEntry:
		return "Second Entry"
	}

	return ""
}

func (m model) blogView() string {
	menuContent := m.getBlogMenuContent()

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		menuContent,
		"  ",
	)
}

