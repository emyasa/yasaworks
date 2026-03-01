package blog

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) navView() string {
	accent := m.Theme.TextAccent().Render
	base := m.Theme.Base().Render

	var navParts []string

	pinnedEntry := pinnedEntries[m.selectedEntryIndex]
	if pinnedEntry.pageIndex > 0 {
		prevNav := base("<< ") + accent("N") + base(" prev")
		navParts = append(navParts, prevNav)
	}

	if pinnedEntry.pageIndex < pinnedEntry.totalPages(m.entryHeight)-1 {
		nextNav := accent("n") + base(" next >>")
		navParts = append(navParts, nextNav)
	}

	nav := strings.Join(navParts, " | ")

	return m.Theme.Base().
		Width(m.navWidth).
		Align(lipgloss.Right).
		Render(nav)
}
