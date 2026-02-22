package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/emyasa/yasaworks/internal/tui/chat"
)

func (m *model) hotKeyUpdate(msg tea.Msg) (tea.Cmd, bool) {
	mode := m.chat.Mode
	if mode == chat.Insert {
		return nil, false
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "l":
			m.page = blogPage
		case "m":
			if m.page == termsPage {
				return nil, false
			}

			m.page = termsPage
			return m.terms.Init(), true
		case "p":
			m.page = chatPage
			return m.chat.Init(), true
		case "q":
			return tea.Quit, true
		}
	}

	return nil, false
}


