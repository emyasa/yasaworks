package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/emyasa/yasaworks/internal/db"
	"github.com/emyasa/yasaworks/internal/tui"
)

func main() {
	database := db.New()
	defer database.Close()

	model, err := tui.NewModel(database, "fingerprint", false, nil)
	if err != nil {
		panic(err)
	}

	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	}
}

