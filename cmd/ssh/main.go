package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/emyasa/yasaworks/internal/tui"
	"github.com/google/uuid"
	gossh "golang.org/x/crypto/ssh"
)

func main() {
	s, err := wish.NewServer(
		wish.WithAddress(":22"),
		wish.WithHostKeyPath(".ssh/host_key"),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			activeterm.Middleware(),
		),
		wish.WithPublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
			hash := md5.Sum(key.Marshal())
			fingerprint := hex.EncodeToString(hash[:])
			ctx.SetValue("fingerprint", fingerprint)
			ctx.SetValue("anonymous", false)
			return true
		}),
		wish.WithKeyboardInteractiveAuth(
			func(ctx ssh.Context, challenger gossh.KeyboardInteractiveChallenge) bool {
				ctx.SetValue("fingerprint", uuid.NewString())
				ctx.SetValue("anonymous", true)
				return true
			},
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("SSH TUI running on port 22")
	log.Fatal(s.ListenAndServe())
}

func teaHandler (s ssh.Session) (tea.Model, []tea.ProgramOption) {
	fingerprint := s.Context().Value("fingerprint").(string)
	anonymous := s.Context().Value("anonymous").(bool)

	clientAddress := s.RemoteAddr().String()
	host, _, _ := net.SplitHostPort(clientAddress)

	model, err := tui.NewModel(fingerprint, anonymous, &host)
	if err != nil {
		return nil, []tea.ProgramOption{}
	}

	return model, []tea.ProgramOption{tea.WithAltScreen()}
}

