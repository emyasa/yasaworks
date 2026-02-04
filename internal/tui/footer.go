package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) footerView() string {
	table := lipgloss.NewStyle().
		Width(78).
		BorderTop(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.AdaptiveColor{
			Dark: "#3A3F42",
			Light: "#D7DBDF",
		}).
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

