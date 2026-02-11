package blog

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) navView(entries []blogEntry, selected int) string {
	accent := m.Theme.TextAccent().Render
	base := m.Theme.Base().Render

	var navParts []string

	vp := entries[selected].viewport
	if vp.YOffset > 0 {
		prevNav := base("<< ") + accent("N") + base(" prev")
		navParts = append(navParts, prevNav)
	}

	if vp.YOffset + vp.Height < vp.TotalLineCount() {
		nextNav := accent("n") + base(" next >>")
		navParts = append(navParts, nextNav)
	}

	nav := strings.Join(navParts, " | ")

	return m.Theme.Base().
		Width(60).
		Align(lipgloss.Right).
		Render(nav)
}

