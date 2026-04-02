// Package admin
package admin

import "github.com/emyasa/yasaworks/internal/tui/theme"

type Model struct {}

func NewModel(theme theme.Theme) Model {
	return Model{}
}

func (m Model) View() string {
	return "admin"
}

