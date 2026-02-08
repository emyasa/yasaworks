package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) footerView() string {
	table := lipgloss.NewStyle().
		Width(78).
		BorderTop(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(m.theme.Border()).
		PaddingBottom(1).
		Align(lipgloss.Center).
		Render()

	return lipgloss.Place(
		m.widthContainer,
		lipgloss.Height(table),
		lipgloss.Center,
		lipgloss.Center,
		table,
	)
}

