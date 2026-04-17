// Package chat handles conversation between potential client and the company
package chat

import (
	"context"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/emyasa/yasaworks/internal/db"
	"github.com/emyasa/yasaworks/internal/registry"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type mode = int
const (
	Normal mode = iota
	Insert
)

type message struct {
	content string
	timestamp time.Time
	isFromSender bool
}

type Model struct {
	ctx context.Context
	db *db.DB
	theme theme.Theme
	conn *registry.Connection
	input textinput.Model
	Mode mode
	messages []message
}

func NewModel(ctx context.Context, db *db.DB, theme theme.Theme, conn *registry.Connection) Model {
	messages, _ := db.ListMessages(ctx, conn.Fingerprint)

	ti := textinput.New()
	ti.Prompt = "> "
	ti.Placeholder = "type a message..."

	// Remove default styling for clean terminal look
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))

	return Model{
		ctx: ctx,
		db: db,
		theme: theme,
		conn: conn,
		input: ti,
		messages: mapMessages(messages),
	}
}

type adminMessageEvent registry.MessageEvent

func readFromChannel(conn *registry.Connection) tea.Cmd {
	if conn == nil {
		return nil
	}

	return func() tea.Msg {
		return adminMessageEvent(conn.FetchMessage())
	}
}

func (m *Model) Init() tea.Cmd {
	m.Mode = Insert
	m.input.Focus()
	return tea.Batch(m.input.Cursor.BlinkCmd(), readFromChannel(m.conn))
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case adminMessageEvent:
		aMsgEvent := registry.MessageEvent(msg)
		m.updateChats(aMsgEvent.Message, false)

		return m, readFromChannel(m.conn)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			if m.Mode == Insert {
				m.Mode = Normal
				m.input.Cursor.SetMode(cursor.CursorStatic)
			}
		case "i":
			if m.Mode == Normal {
				m.Mode = Insert
				m.input.Cursor.SetMode(cursor.CursorBlink)

				return m, m.input.Cursor.BlinkCmd()
			}
		case "enter":
			if text := m.input.Value(); text != "" {
				m.updateChats(text, true)
				m.input.SetValue("")

				registry.HandleClientMessage(m.conn, text)

				createMessageRequest := db.CreateMessageRequest{
					ClientFingerprint: m.conn.Fingerprint,
					SenderType: db.SenderClient,
					Content: text,
				}
				m.db.CreateMessage(m.ctx, createMessageRequest)
			}
		}
	}

	if m.Mode == Normal {
		return m, nil
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	return m, cmd
}

func (m *Model) updateChats(content string, isSender bool) {
	message := message{
		content: content,
		timestamp: time.Now(),
		isFromSender: isSender,
	}

	m.messages = append(m.messages, message)
}

func (m Model) View() string {
	sb := strings.Builder{}
	for _, message := range m.messages {
		bubbleStyle := m.theme.ReceiverBubbleStyle()
		position := lipgloss.Left

		if message.isFromSender {
			bubbleStyle = m.theme.SenderBubbleStyle()
			position = lipgloss.Right
		}

		timestamp := message.timestamp.Format("15:04")
		timestampView := m.theme.TimestampStyle().Render(timestamp)

		messageView := lipgloss.Place(80, 1, position, lipgloss.Bottom, bubbleStyle.Render(message.content) + timestampView)
		sb.WriteString(messageView + "\n")
	}
	messagesView := sb.String()

	inputView := m.theme.Base().
		MarginLeft(1).
		Render(m.input.View())

	modeString := ""
	if m.Mode == Insert {
		modeString = "-- INSERT --"
	}

	statusLine := m.theme.Base().
		MarginLeft(1).
		Render(modeString)

	child := lipgloss.JoinVertical(lipgloss.Left, messagesView, inputView, statusLine)

	return lipgloss.Place(80, 24, lipgloss.Left, lipgloss.Bottom, child)
}

