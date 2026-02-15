package blog

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type blogEntry struct {
	name string
	mdPath string
	content string
	lines []string
	pageIndex int
}

func maxEntryWidth(entries []*blogEntry) int {
	max := 0
	for _, e := range entries {
		if w := lipgloss.Width(e.name); w > max {
			max = w
		}
	}

	return max
}

func (b blogEntry) totalPages(pageHeight int) int {
	totalLines := len(b.lines)
	return (totalLines + pageHeight - 1) / pageHeight
}

func (b blogEntry) visibleContent(pageHeight int) string {
	start := b.pageIndex * pageHeight
	end := start + pageHeight
	end = min(end, len(b.lines))

	return strings.Join(b.lines[start:end], "\n")
}

func (m Model) entryView() string {
	entry := m.blogEntries[m.selected]
	entryVisibleContent := entry.visibleContent(m.contentHeight)
	content := lipgloss.JoinVertical(
		lipgloss.Top,
		entryVisibleContent,
		" ",
		m.navView(),
	)

	return m.Theme.Base().
		MarginTop(1).
		Render(content)
}

