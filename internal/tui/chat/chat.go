// Package chat handles conversation between potential client and the company
package chat

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/emyasa/yasaworks/internal/db"
	rl "github.com/emyasa/yasaworks/internal/ratelimiter"
	"github.com/emyasa/yasaworks/internal/registry"
	"github.com/emyasa/yasaworks/internal/tui/theme"
)

type mode = int
const (
	Normal mode = iota
	Insert
)

const (
	messagesBufferSize = 80 //buffer size must be at least x4 of window size
	messagesWindowSize = 18
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
	messagesCursorIndex int
	messages []message
	hasReachedStart bool
	hasReachedEnd bool
	modeString string
}

func NewModel(ctx context.Context, db *db.DB, theme theme.Theme, conn *registry.Connection) Model {
	if db == nil {
		log.Fatal("chat's NewModel requires *db")
	}

	if conn == nil {
		log.Fatal("chat's NewModel requires *conn")
	}

	messages, err := db.ListMessages(ctx, conn.Fingerprint, messagesBufferSize, nil)
	if err != nil {
		log.Fatalf("chat's NewModel error: %s", err)
	}

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
		messagesCursorIndex: len(messages) - 1,
		messages: mapMessages(messages),
		hasReachedStart: len(messages) <= messagesWindowSize || len(messages) <= messagesBufferSize,
		hasReachedEnd: true,
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
				m.modeString = ""
				m.Mode = Normal
				m.input.Cursor.SetMode(cursor.CursorStatic)
			}
		case "i":
			if m.Mode == Normal {
				m.modeString = "-- INSERT --"
				m.Mode = Insert
				m.input.Cursor.SetMode(cursor.CursorBlink)

				return m, m.input.Cursor.BlinkCmd()
			}
		case "k":
			if m.canScrollUp() {
				m.scrollUp()
			}
		case "j":
			if m.canScrollDown() {
				m.scrollDown()
			}
		case "enter":
			allow := rl.Allow(m.conn.Fingerprint)
			if text := m.input.Value(); text != "" && allow {
				m.updateChats(text, true)
				m.input.SetValue("")

				registry.HandleClientMessage(m.conn, text)

				createMessageRequest := db.CreateMessageRequest{
					ClientFingerprint: m.conn.Fingerprint,
					SenderType: db.SenderClient,
					Content: text,
				}
				m.db.CreateMessage(m.ctx, createMessageRequest)
			} else {
				m.modeString = "Rate Limit Reached."
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
	m.messagesCursorIndex++
}

func (m Model) View() string {
	windowStart := max(0, m.messagesCursorIndex - messagesWindowSize + 1)
	messages := m.messages[windowStart : m.messagesCursorIndex + 1]

	sb := strings.Builder{}
	var lastDate string

	for _, message := range messages {
		currentDate := message.timestamp.Format("2006-01-02")
		if lastDate != currentDate {
			dateLabel := message.timestamp.Format("Jan 02, 2006")
			separator := lipgloss.NewStyle().
				Width(80).
				Align(lipgloss.Center).
				Foreground(m.theme.Border()).
				Render("---- " + dateLabel + " -----")

			sb.WriteString(separator + "\n")
		}
		lastDate = currentDate

		bubbleStyle := m.theme.ReceiverBubbleStyle()
		position := lipgloss.Left

		if message.isFromSender {
			bubbleStyle = m.theme.SenderBubbleStyle()
			position = lipgloss.Right
		}

		timestamp := message.timestamp.Format("15:04")
		timestampView := m.theme.TimestampStyle().Render(timestamp)

		messageView := lipgloss.PlaceHorizontal(80, position, bubbleStyle.Render(message.content) + timestampView)
		sb.WriteString(messageView + "\n")
	}
	messagesView := sb.String()

	inputView := m.theme.Base().
		MarginLeft(1).
		Render(m.input.View())

	statusLine := m.theme.Base().
		MarginLeft(1).
		Render(m.modeString)

	child := lipgloss.JoinVertical(
		lipgloss.Left,
		messagesView,
		inputView,
		statusLine,
	)

	return lipgloss.Place(80, 24, lipgloss.Left, lipgloss.Bottom, child)
}

