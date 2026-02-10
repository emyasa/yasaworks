// Package blog handles all the blog-related entries
package blog

import (
	"embed"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
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
	Viewport *viewport.Model
}

//go:embed entries/*.md
var entriesFS embed.FS
var blogEntries = []BlogEntry{}

func (m Model) Init() tea.Cmd {
	r, _ := glamour.NewTermRenderer(
	glamour.WithAutoStyle(),
	glamour.WithWordWrap(50))

	firstEntryContent, err := entriesFS.ReadFile("entries/first.md")
	if err != nil {
		panic(err)
	}

	detailContent, _ := r.Render(string(firstEntryContent))
	vp := viewport.New(50, 10)
	vp.SetContent(detailContent)

	blogEntries = append(blogEntries, BlogEntry{
		Name: "First Entry",
		Content: string(firstEntryContent),
		Viewport: &vp,
	})
	
	return nil
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
			vp := blogEntries[m.selected].Viewport
			vp.ScrollDown(10)

			return m, nil
		case "p":
			vp := blogEntries[m.selected].Viewport
			vp.ScrollUp(10)

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
	vp := entries[selected].Viewport
	var navParts []string
	if vp.YOffset > 0 {
		navParts = append(navParts, "<< p prev")
	}

	if vp.YOffset+vp.Height < vp.TotalLineCount() {
		navParts = append(navParts, "n next >>")
	}

	nav := strings.Join(navParts, " | ")
	containerWidth := 80 - m.menuWidth

	navStyle := m.Theme.Base().
		Width(45).
		Align(lipgloss.Right)

	content := lipgloss.JoinVertical(
		lipgloss.Top,
		vp.View(),
		navStyle.Render(nav),
	)

	containerStyle := m.Theme.Base().
		Width(containerWidth).
		MarginTop(1).
		Padding(0, 1)

	return containerStyle.Render(content)
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

