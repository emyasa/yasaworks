package main

import (
	"log"

    tea "github.com/charmbracelet/bubbletea"
	"github.com/emyasa/yasaworks/internal/tui"
)

func main() {
	model, err := tui.NewModel()
	if err != nil {
		panic(err)
	}

	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	}
}

