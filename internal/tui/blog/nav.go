package blog

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) navView(entries []*blogEntry, selected int) string {
	accent := m.Theme.TextAccent().Render
	base := m.Theme.Base().Render

	var navParts []string

	blogEntry := entries[selected]
	if blogEntry.pageIndex > 0 {
		prevNav := base("<< ") + accent("N") + base(" prev")
		navParts = append(navParts, prevNav)
	}

	if blogEntry.pageIndex < blogEntry.totalPages(m.contentHeight) - 1 {
		nextNav := accent("n") + base(" next >>")
		navParts = append(navParts, nextNav)
	}

	nav := strings.Join(navParts, " | ")

	return m.Theme.Base().
		Width(m.navWidth).
		Align(lipgloss.Right).
		Render(nav)
}

