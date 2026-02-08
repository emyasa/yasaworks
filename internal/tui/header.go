package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func (m model) headerView() string {
	bold := m.theme.TextAccent().Bold(true).Render
	accent := m.theme.TextAccent().Render
	base := m.theme.Base().Render

	logo := bold("yasaworks")
	blog := bold("b") + accent(" blog")
	rtfwm := accent("r") + base(" TFWM")
	contact := accent("c") + base(" contact")

	tabs := []string{
		logo,
		blog,
		rtfwm,
		contact,
	}

	table := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(m.theme.Base().Foreground(m.theme.Border())).
		Row(tabs...).
		Width(m.widthContainer).
		StyleFunc(func(row, col int) lipgloss.Style {
			return m.theme.Base().
				Padding(0, 1).
				AlignHorizontal(lipgloss.Center)
		}).
		Render()

	return lipgloss.Place(
		m.widthContainer,
		lipgloss.Height(table),
		lipgloss.Center,
		lipgloss.Center,
		table,
	)
}

