// Package theme handles the shared rendering context of the tui
package theme

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	renderer *lipgloss.Renderer
	base lipgloss.Style

	brand lipgloss.TerminalColor
	accent lipgloss.TerminalColor
}

func BasicTheme() Theme {
	defaultRenderer := lipgloss.DefaultRenderer()
	baseColor := lipgloss.AdaptiveColor{Dark: "#889096", Light: "#889096"}

	brandColor := lipgloss.Color("#FF0000")
	accentColor := lipgloss.AdaptiveColor{Dark: "#FFFFFF", Light: "#11181C"}

	return Theme{
		renderer: defaultRenderer,
		base: defaultRenderer.NewStyle().Foreground(baseColor),
		brand: brandColor,
		accent: accentColor,
	}
}

func (b Theme) Base() lipgloss.Style {
	return b.base
}

func (b Theme) TextAccent() lipgloss.Style {
	return b.Base().Foreground(b.accent)
}

func (b Theme) Brand() lipgloss.TerminalColor {
	return b.brand
}

