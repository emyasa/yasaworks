// Package chat handles conversation between potential client and the company
package chat

import "github.com/emyasa/yasaworks/internal/tui/theme"

type Model struct {
	theme theme.Theme
}

func (m Model) View() string {
	return ""
}

