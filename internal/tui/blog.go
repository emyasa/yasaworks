package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) getBlogMenuContent() string {
	pages := strings.Builder{}

	for _, p := range m.blogEntries {
		name := getBlogEntryName(p)
		pages.WriteString(name + "\n")
	}

	return lipgloss.NewStyle().
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

