package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) footerView() string {
	bold := m.theme.TextAccent().Bold(true).Render
	base := m.theme.Base().Render
	table := lipgloss.NewStyle().
		Width(78).
		BorderTop(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(m.theme.Border()).
		PaddingBottom(1).
		Align(lipgloss.Right)

	navCommands := []string{}
	navCommands = append(navCommands, bold("j") + " " + base("↓"))
	navCommands = append(navCommands, bold("k") + " " + base("↑") + " ")
	footer := table.Render(strings.Join(navCommands, " "))

	return lipgloss.Place(
		m.widthContainer,
		lipgloss.Height(footer),
		lipgloss.Center,
		lipgloss.Center,
		footer,
	)
}

