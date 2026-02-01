package splash

import "github.com/charmbracelet/lipgloss"

func(m Model) logoView() string {
	return lipgloss.NewStyle().Bold(true).
		Render("yasaworks") + m.cursorView()
}

