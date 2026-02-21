package tui

import "github.com/charmbracelet/lipgloss"

func (m model) layout(header, content, footer string) string {
		height := m.heightContainer
		height -= lipgloss.Height(header)
		height -= lipgloss.Height(footer)

		body := lipgloss.NewStyle().
			Width(80).
			Height(height).
			Render(content)

		child := lipgloss.JoinVertical(
			lipgloss.Left,
			header,
			body,
			footer,
		)

		return lipgloss.Place(
			m.viewportWidth,
			m.viewportHeight,
			lipgloss.Center,
			lipgloss.Center,
			m.theme.Base().
				MaxWidth(m.widthContainer).
				MaxHeight(m.heightContainer).
				Render(child),
		)
}

