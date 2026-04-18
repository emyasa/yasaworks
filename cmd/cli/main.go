package main

import (
	"context"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/emyasa/yasaworks/internal/db"
	"github.com/emyasa/yasaworks/internal/tui"
)

func main() {
	db := db.New()
	defer db.Close()

	model, err := tui.NewModel(context.Background(), db, false, nil)
	if err != nil {
		panic(err)
	}

	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	}
}

