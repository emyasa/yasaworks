// Package blog handles all the blog-related entries
package blog

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	selected int
}

type blogEntry struct {
	name string
	content viewport.Model
}

var blogEntries = []blogEntry{
	{name: "First Entry"},
	{name: "Next Entry"},
}

func (m Model) BlogView() string {
	menuContent := renderBlogMenu(blogEntries, m.selected)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		menuContent,
		"  ",
	)
}

func renderBlogMenu(entries []blogEntry, selected int) string {
	menuWidth := maxEntryWidth(entries)

	var sb strings.Builder
	for i, e := range entries {
		menuItemStyle := lipgloss.NewStyle().
			Width(menuWidth + 2).
			Padding(0, 1)

		if i == selected {
			menuItemStyle = menuItemStyle.Background(lipgloss.Color("#555")).
				Foreground(lipgloss.Color("#fff")).
				Bold(true)
		}

		sb.WriteString(menuItemStyle.Render(e.name))
		if i < len(entries) - 1 {
			sb.WriteString("\n")
		}
	}

	containerStyle := lipgloss.NewStyle().
		MarginTop(1).
		Padding(0, 1)
	
	return containerStyle.Render(sb.String())
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

