// Package splash handles the tui's splash state
package splash

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	state State
}

type State struct {
	cursor cursorState
}

func(m Model) SplashInit() tea.Cmd {
	return m.cursorInit()
}

