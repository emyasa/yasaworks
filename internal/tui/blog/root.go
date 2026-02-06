// Package blog handles all the blog-related entries
package blog

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type blogEntry = int
const (
	firstEntry blogEntry = iota
	secondEntry
)

var blogEntries = []blogEntry{
	firstEntry,
	secondEntry,
}

func BlogView() string {
	menuContent := getBlogMenuContent()

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		menuContent,
		"  ",
	)
}

func getBlogMenuContent() string {
	menuWidth := 0
	pages := strings.Builder{}

	for _, p := range blogEntries {
		w := lipgloss.Width(getBlogEntryName(p))
		if w > menuWidth {
			menuWidth = w
		}
	}

	menuItem := lipgloss.NewStyle().
		Width(menuWidth+2).
		Padding(0, 1)

	for _, p := range blogEntries {
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

