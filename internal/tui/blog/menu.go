package blog

import "strings"

func (m Model) menuView() string {
	var sb strings.Builder

	headerStyle := m.Theme.Base().
		Width(m.menuWidth+2).
		Padding(0, 1).
		Bold(true).
		Foreground(m.Theme.Accent())
	sb.WriteString(headerStyle.Render("Pinned"))
	sb.WriteString("\n")

	for i, e := range pinnedEntries {
		menuItemStyle := m.Theme.Base().
			Width(m.menuWidth+2).
			Padding(0, 1)

		if i == m.selectedEntryIndex {
			menuItemStyle = menuItemStyle.Background(m.Theme.Highlight()).
				Foreground(m.Theme.Accent()).
				Bold(true)
		}

		sb.WriteString(menuItemStyle.Render(e.name))
		if i < len(pinnedEntries)-1 {
			sb.WriteString("\n")
		}
	}

	containerStyle := m.Theme.Base().
		MarginTop(1).
		Padding(0, 1)

	return containerStyle.Render(sb.String())
}
