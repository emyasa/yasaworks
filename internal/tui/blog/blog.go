// Package blog handles all the blog-related entries
package blog

import (
	"embed"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type Model struct {
	Theme theme.Theme
	ContainerWidth int
	ContainerHeight int

	menuWidth int
	contentWidth int
	contentHeight int
	navWidth int
	selected int
}

//go:embed entries/*.md
var entriesFS embed.FS
var blogEntries = []*blogEntry{
	{name: "Dev Workflow Journey", mdPath: "entries/dev-workflow.md"},
}

//go:embed styles/dark.json
var darkStyle []byte

func NewModel(theme theme.Theme, containerWidth int, containerHeight int) Model {
	menuWidth := maxEntryWidth(blogEntries) + 6
	contentWidth := containerWidth - menuWidth
	navWidth := contentWidth - 6 
	pageHeight := containerHeight - 10

	r, _ := glamour.NewTermRenderer(
		glamour.WithStylesFromJSONBytes(darkStyle),
		glamour.WithWordWrap(contentWidth))

	for i, entry := range blogEntries {
		content, err := entriesFS.ReadFile(entry.mdPath)
		if err != nil {
			panic(err)
		}

		detailContent, _ := r.Render(string(content))
		blogEntries[i].content = detailContent
		blogEntries[i].lines = strings.Split(detailContent, "\n")
	}
	
	return Model{
		Theme: theme,
		menuWidth: menuWidth,
		contentWidth: contentWidth,
		contentHeight: pageHeight,
		navWidth: navWidth,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "j", "down":
			if m.selected < len(blogEntries) - 1 {
				m.selected += 1
			}

			return m, nil
		case "shift+tab", "k", "up":
			if m.selected > 0 {
				m.selected -= 1
			}
			
			return m, nil
		case "n":
			blogEntry := blogEntries[m.selected]
			if blogEntry.pageIndex < blogEntry.totalPages(m.contentHeight) - 1 {
				blogEntry.pageIndex++
			}

			return m, nil
		case "N":
			blogEntry := blogEntries[m.selected]
			if blogEntry.pageIndex > 0 {
				blogEntry.pageIndex--
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

func (m Model) renderBlogMenu(entries []*blogEntry, selected int) string {
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

		sb.WriteString(menuItemStyle.Render(e.name))
		if i < len(entries) - 1 {
			sb.WriteString("\n")
		}
	}

	containerStyle := m.Theme.Base().
		MarginTop(1).
		Padding(0, 1)

	return containerStyle.Render(sb.String())
}

func (m Model) renderBlogDetail(entries []*blogEntry, selected int) string {
	entryVisibleContent := entries[selected].visibleContent(m.contentHeight)
	content := lipgloss.JoinVertical(
		lipgloss.Top,
		entryVisibleContent,
		" ",
		m.navView(entries, selected),
	)

	return m.Theme.Base().
		MarginTop(1).
		Render(content)
}

