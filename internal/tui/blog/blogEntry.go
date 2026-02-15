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

func (b blogEntry) totalPages(pageHeight int) int {
	totalLines := len(b.lines)
	return (totalLines + pageHeight - 1) / pageHeight
}

func (b blogEntry) visibleContent(pageHeight int) string {
	start := b.pageIndex * pageHeight
	end := start + pageHeight

	if end > len(b.lines) {
		end = len(b.lines)
	}

	return strings.Join(b.lines[start:end], "\n")
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

