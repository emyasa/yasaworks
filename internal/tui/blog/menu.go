package blog

import "strings"

func (m Model) constructEntries(sb *strings.Builder, predicate func(*blogEntry) bool) {
	for i, e := range blogEntries {
		if !predicate(e) {
			continue
		}

		menuItemStyle := m.Theme.Base().
			Width(m.menuWidth+2).
			Padding(0, 0, 0, 1)

		if i == m.selectedEntryIndex {
			menuItemStyle = menuItemStyle.Background(m.Theme.Highlight()).
				Foreground(m.Theme.Accent()).
				Bold(true)
		}

		sb.WriteString(menuItemStyle.Render(e.name))
		if i < len(blogEntries)-1 {
			sb.WriteString("\n")
		}
	}
}

func (m Model) menuView() string {
	headerStyle := m.Theme.Base().
		Width(m.menuWidth+2).
		Bold(true).
		Padding(0, 1)

	var sb strings.Builder
	sb.WriteString(headerStyle.Render("Pinned"))
	sb.WriteString("\n")

	m.constructEntries(&sb, func(e *blogEntry) bool { return e.pinned })
	sb.WriteString("\n")

	m.constructEntries(&sb, func(e *blogEntry) bool { return !e.pinned })

	containerStyle := m.Theme.Base().
		MarginTop(1).
		Padding(0, 1)

	return containerStyle.Render(sb.String())
}

