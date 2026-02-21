package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/emyasa/yasaworks/internal/tui/chat"
)

func (m *model) hotKeyUpdate(msg tea.Msg) tea.Cmd {
	mode := m.chat.Mode
	if mode == chat.Insert {
		return nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "l":
			m.page = blogPage
		case "m":
			if m.page == termsPage {
				return nil
			}

			m.page = termsPage
			return m.terms.Init()
		case "p":
			m.page = chatPage
			m.chat.Init()
		case "q":
			return tea.Quit
		}
	}

	return nil
}


