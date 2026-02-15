package blog

import "strings"

func (m Model) menuView() string {
	var sb strings.Builder
	for i, e := range blogEntries {
		menuItemStyle := m.Theme.Base().
			Width(m.menuWidth + 2).
			Padding(0, 1)

		if i == m.selectedEntryIndex {
			menuItemStyle = menuItemStyle.Background(m.Theme.Highlight()).
				Foreground(m.Theme.Accent()).
				Bold(true)
		}

		sb.WriteString(menuItemStyle.Render(e.name))
		if i < len(blogEntries) - 1 {
			sb.WriteString("\n")
		}
	}

	containerStyle := m.Theme.Base().
		MarginTop(1).
		Padding(0, 1)

	return containerStyle.Render(sb.String())
}
 
