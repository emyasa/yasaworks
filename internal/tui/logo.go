package tui

import "github.com/charmbracelet/lipgloss"

func(m model) LogoView() string {
	return lipgloss.NewStyle().Bold(true).
		Render("yasaworks") + m.CursorView()
}

