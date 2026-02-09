// Package blog handles all the blog-related entries
package blog

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type Model struct {
	Theme theme.Theme
	menuWidth int
	selected int
}

type BlogEntry struct {
	Name string
	Content string
}

var blogEntries = []BlogEntry{}

func Register(blogEntry BlogEntry) {
	blogEntries = append(blogEntries, blogEntry)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "down", "j":
			if m.selected < len(blogEntries) - 1 {
				m.selected += 1
			}

			return m, nil
		case "shift+tab", "up", "k":
			if m.selected > 0 {
				m.selected -= 1
			}
			
			return m, nil
		}
	}

	return m, nil
}

func (m Model) View() string {
	menuContent := m.renderBlogMenu(blogEntries, m.selected)
	detailContent := m.renderBlogDetail(blogEntries, m.selected)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		menuContent,
		"  ",
		detailContent,
	)
}

func (m Model) renderBlogMenu(entries []BlogEntry, selected int) string {
	m.menuWidth = maxEntryWidth(entries)

	var sb strings.Builder
	for i, e := range entries {
		menuItemStyle := m.Theme.Base().
			Width(m.menuWidth + 2).
			Padding(0, 1)

		if i == selected {
			menuItemStyle = menuItemStyle.Background(m.Theme.Highlight()).
				Foreground(m.Theme.Accent()).
				Bold(true)
		}

		sb.WriteString(menuItemStyle.Render(e.Name))
		if i < len(entries) - 1 {
			sb.WriteString("\n")
		}
	}

	containerStyle := m.Theme.Base().
		MarginTop(1).
		Padding(0, 1)

	return containerStyle.Render(sb.String())
}

func (m Model) renderBlogDetail(entries []BlogEntry, selected int) string {
	containerStyle := m.Theme.Base().
		Width(80 - m.menuWidth).
		MarginTop(1).
		Padding(0, 1)
	
	return containerStyle.Render(entries[selected].Content)
}

func maxEntryWidth(entries []BlogEntry) int {
	max := 0
	for _, e := range entries {
		if w := lipgloss.Width(e.Name); w > max {
			max = w
		}
	}

	return max
}

