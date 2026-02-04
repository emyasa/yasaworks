package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func (m model) HeaderView() string {
	logo := lipgloss.NewStyle().Bold(true).Render("yasaworks")
	blog := lipgloss.NewStyle().Render("blog")
	rtfwm := lipgloss.NewStyle().Render("FWM")
	contact := lipgloss.NewStyle().Render("contact")

	tabs := []string{
		logo,
		blog,
		rtfwm,
		contact,
	}

	table := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
			Dark: "#3A3F42",
			Light: "#D7DBDF",
		})).
		Rows(tabs).
		Width(78).
		StyleFunc(func(row, col int) lipgloss.Style {
			return lipgloss.NewStyle().
				Padding(0, 1).
				AlignHorizontal(lipgloss.Center)
		}).
		Render()

	return lipgloss.Place(
		m.viewportWidth,
		lipgloss.Height(table),
		lipgloss.Center,
		lipgloss.Center,
		table,
	)
}

