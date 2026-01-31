package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/emyasa/yasaworks/internal/tui"
)

func main() {
	s, err := wish.NewServer(
		wish.WithAddress(":22"),
		wish.WithHostKeyPath(".ssh/host_key"),
		wish.WithMiddleware(
			bubbletea.Middleware(func(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
				model, err := tui.NewModel()
				if err != nil {
					return nil, []tea.ProgramOption{}
				}
				
				return model, []tea.ProgramOption{tea.WithAltScreen()}
			}),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("SSH TUI running on port 22")
	log.Fatal(s.ListenAndServe())
}

