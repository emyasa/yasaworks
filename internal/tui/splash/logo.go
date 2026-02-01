package splash

import "github.com/charmbracelet/lipgloss"

func(m Model) LogoView() string {
	return lipgloss.NewStyle().Bold(true).
		Render("yasaworks") + m.CursorView()
}

