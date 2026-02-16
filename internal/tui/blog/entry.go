package blog

import (
	"embed"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type blogEntry struct {
	name string
	mdPath string
	content string
	lines []string
	pageIndex int
}

//go:embed entries/*.md
var entriesFS embed.FS
var blogEntries = []*blogEntry{
	{name: "Dev Workflow Journey", mdPath: "entries/dev-workflow.md"},
}

func setupEntries(entryWidth int, markdownStyle glamour.TermRendererOption) {
	r, _ := glamour.NewTermRenderer(
		markdownStyle,
		glamour.WithWordWrap(entryWidth))

	for i, entry := range blogEntries {
		content, err := entriesFS.ReadFile(entry.mdPath)
		if err != nil {
			panic(err)
		}

		detailContent, _ := r.Render(string(content))
		blogEntries[i].content = detailContent
		blogEntries[i].lines = strings.Split(detailContent, "\n")
	}
}

func maxEntryWidth() int {
	max := 0
	for _, e := range blogEntries {
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

func (m *Model) getNextEntry() {
	if m.selectedEntryIndex < len(blogEntries) - 1 {
		m.selectedEntryIndex += 1
	}
}

func (m *Model) getPrevEntry() {
	if m.selectedEntryIndex > 0 {
		m.selectedEntryIndex -= 1
	}
}

func (m Model) entryNextPage() {
	entry := blogEntries[m.selectedEntryIndex]
	if entry.pageIndex < entry.totalPages(m.entryHeight) - 1 {
		entry.pageIndex++
	}
}

func (m Model) entryPrevPage() {
	entry := blogEntries[m.selectedEntryIndex]
	if entry.pageIndex > 0 {
		entry.pageIndex--
	}
}

func (m Model) entryView() string {
	entry := blogEntries[m.selectedEntryIndex]
	entryVisibleContent := entry.visibleContent(m.entryHeight)
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

