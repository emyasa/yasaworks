package admin

import "github.com/charmbracelet/lipgloss"

func (m Model) chatPanelView() string {
	inputView := m.theme.Base().
		MarginLeft(1).
		Render(m.input.View())

	modeString := ""
	if m.Mode == Insert {
		modeString = "-- INSERT --"
	}

	statusLine := m.theme.Base().
		MarginLeft(1).
		Render(modeString)

	child := lipgloss.JoinVertical(lipgloss.Left, inputView, statusLine)

	return lipgloss.Place(80, 23, lipgloss.Left, lipgloss.Bottom, child)
}

