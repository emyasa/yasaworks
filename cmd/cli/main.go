package main

import (
	"context"
	"io"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/emyasa/yasaworks/internal/ctxkeys"
	"github.com/emyasa/yasaworks/internal/db"
	"github.com/emyasa/yasaworks/internal/registry"
	"github.com/emyasa/yasaworks/internal/tui"
)

func main() {
	log.SetOutput(io.Discard)

	db := db.New()
	defer db.Close()

	ctx := context.WithValue(context.Background(), ctxkeys.IsAdmin, false)
	conn := &registry.Connection{Fingerprint: "4e4aa06501547c64cbbe41c5fa7a7b67"}
	model, err := tui.NewModel(ctx, db, conn)
	if err != nil {
		panic(err)
	}

	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	}
}

