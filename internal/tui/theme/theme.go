// Package theme handles the shared rendering context of the tui
package theme

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	base lipgloss.Style

	brand lipgloss.TerminalColor
	accent lipgloss.TerminalColor
	border lipgloss.TerminalColor
	highlight lipgloss.TerminalColor
}

func BasicTheme() Theme {
	defaultRenderer := lipgloss.DefaultRenderer()
	baseColor := lipgloss.AdaptiveColor{Dark: "#889096", Light: "#889096"}

	brandColor := lipgloss.Color("#FF0000")
	accentColor := lipgloss.AdaptiveColor{Dark: "#FFFFFF", Light: "#11181C"}
	borderColor := lipgloss.AdaptiveColor{Dark: "#53565A", Light: "#D7DBDF"}
	highlight := brandColor

	return Theme{
		base: defaultRenderer.NewStyle().Foreground(baseColor),
		brand: brandColor,
		accent: accentColor,
		border: borderColor,
		highlight: highlight,
	}
}

func (b Theme) Base() lipgloss.Style {
	return b.base
}

func (b Theme) TextAccent() lipgloss.Style {
	return b.Base().Foreground(b.accent)
}

func (b Theme) Accent() lipgloss.TerminalColor {
	return b.accent
}

func (b Theme) Brand() lipgloss.TerminalColor {
	return b.brand
}

func (b Theme) Border() lipgloss.TerminalColor {
	return b.border
}

func (b Theme) Highlight() lipgloss.TerminalColor {
	return b.highlight
}

