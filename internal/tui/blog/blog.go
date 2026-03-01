// Package blog handles all the blog-related entries
package blog

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type Model struct {
	Theme theme.Theme

	menuWidth int
	entryWidth int
	entryHeight int
	navWidth int

	selectedEntryIndex int
}

func NewModel(theme theme.Theme, containerWidth int, containerHeight int) Model {
	menuWidth := maxEntryWidth()
	entryWidth := containerWidth - (menuWidth + 6)
	entryHeight := containerHeight - 10
	navWidth := entryWidth - 6 

	setupEntries(entryWidth, theme.MarkdownStyle())
	
	return Model{
		Theme: theme,
		menuWidth: menuWidth,
		entryWidth: entryWidth,
		entryHeight: entryHeight,
		navWidth: navWidth,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "j", "down":
			m.getNextEntry()

			return m, nil
		case "shift+tab", "k", "up":
			m.getPrevEntry()
			
			return m, nil
		case "n":
			m.entryNextPage()

			return m, nil
		case "N":
			m.entryPrevPage()

			return m, nil
		}
	}

	return m, nil
}

func (m Model) View() string {
	menuContent := m.menuView()
	detailContent := m.entryView()

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		menuContent,
		"  ",
		detailContent,
	)
}

