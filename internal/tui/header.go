package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func (m *model) headerUpdate(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "l":
			m.page = blogPage
		case "m":
			m.page = termsPage
			return m.terms.Init()
		case "p":
			m.page = chatPage
		}
	}

	return nil
}

func (m model) headerView() string {
	bold := m.theme.TextAccent().Bold(true).Render
	accent := m.theme.TextAccent().Render
	base := m.theme.Base().Render

	logo := bold("yasaworks")
	blog := accent("l") + base(" logs")
	terms := accent("m") + base(" man interface")
	chat := accent("p") + base(" ping")

	switch m.page {
	case blogPage:
		blog = accent("l logs")
	case termsPage:
		terms = accent("m man interface")
	case chatPage:
		chat = accent("p ping")
	}

	tabs := []string{
		logo,
		blog,
		terms,
		chat,
	}

	table := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(m.theme.Base().Foreground(m.theme.Border())).
		Row(tabs...).
		Width(m.widthContainer - 2).
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

