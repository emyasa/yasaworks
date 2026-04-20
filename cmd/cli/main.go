package main

import (
	"context"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/emyasa/yasaworks/internal/ctxkeys"
	"github.com/emyasa/yasaworks/internal/db"
	"github.com/emyasa/yasaworks/internal/registry"
	"github.com/emyasa/yasaworks/internal/tui"
)

func main() {
	db := db.New()
	defer db.Close()

	ctx := context.WithValue(context.Background(), ctxkeys.IsAdmin, true)
	conn := &registry.Connection{}
	model, err := tui.NewModel(ctx, db, conn)
	if err != nil {
		panic(err)
	}

	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	}
}

