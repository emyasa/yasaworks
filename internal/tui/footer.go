package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const footerWidth = 78

func (m model) footerView() string {
	bold := m.theme.TextAccent().Bold(true).Render
	base := m.theme.Base().Render
	table := lipgloss.NewStyle().
		Width(footerWidth).
		BorderTop(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(m.theme.Border()).
		PaddingBottom(1)

	sb := strings.Builder{}
	navCommandsContent := base("navigate:") + " " + 
		bold("j") + " " + base("↓") +  " " + 
		bold("k") + " " + base("↑")
	widthTaken := lipgloss.Width(navCommandsContent)
	navCommands := lipgloss.NewStyle().
		Width(widthTaken).
		Align(lipgloss.Left).
		Render(navCommandsContent)
	sb.WriteString(navCommands)

	quitCommand := lipgloss.NewStyle().
		Width(footerWidth - widthTaken).
		Align(lipgloss.Right).
		Render(bold("q") + " " + base("quit"))
	sb.WriteString(quitCommand)

	footer := table.Render(sb.String())

	return lipgloss.Place(
		m.widthContainer,
		lipgloss.Height(footer),
		lipgloss.Center,
		lipgloss.Center,
		footer,
	)
}

