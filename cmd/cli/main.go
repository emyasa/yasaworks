package main

import (
	"log"

    tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if _, err := tea.NewProgram(nil, tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	}
}

