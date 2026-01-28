package main

import (
    "log"

    tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/ssh"
)

type model struct{}

func (m model) Init() tea.Cmd { return nil }
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg.(type) {
    case tea.KeyMsg:
        return m, tea.Quit
    }
    return m, nil
}
func (m model) View() string {
    return "Nothing to see here.\n\nPress any key to quit.\n"
}

func main() {
	s, err := wish.NewServer(
		wish.WithAddress(":22"),
		wish.WithHostKeyPath(".ssh/host_key"),
		wish.WithMiddleware(
			bubbletea.Middleware(func(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
				return model{}, []tea.ProgramOption{
					tea.WithAltScreen(),
				}
			}),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("SSH TUI running on port 22")
	log.Fatal(s.ListenAndServe())
}

