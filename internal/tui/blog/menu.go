package blog

import "strings"

func (m Model) constructPinnedEntries(sb *strings.Builder) {
	headerStyle := m.Theme.Base().
		Width(m.menuWidth+2).
		Bold(true).
		Padding(0, 1)

	sb.WriteString(headerStyle.Render("Pinned"))
	sb.WriteString("\n")

	for i, e := range blogEntries {
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
	var sb strings.Builder
	m.constructPinnedEntries(&sb)
	sb.WriteString("\n\n")
	m.constructPinnedEntries(&sb)

	containerStyle := m.Theme.Base().
		MarginTop(1).
		Padding(0, 1)

	return containerStyle.Render(sb.String())
}

