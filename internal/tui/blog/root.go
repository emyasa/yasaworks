// Package blog handles all the blog-related entries
package blog

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type blogEntry struct {
	name string
}

var blogEntries = []blogEntry{
	{name: "First Entry"},
	{name: "Next Entry"},
}

func BlogView() string {
	menuContent := renderBlogMenu(blogEntries)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		menuContent,
		"  ",
	)
}

func renderBlogMenu(entries []blogEntry) string {
	menuWidth := maxEntryWidth(entries)

	menuItemStyle := lipgloss.NewStyle().
		Width(menuWidth + 2).
		Padding(0, 1)
	
	var sb strings.Builder
	for i, e := range entries {
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

