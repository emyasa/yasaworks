// Package admin
package admin

import (
	"context"
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

const ConversationsViewWidth = 17

type message struct {
	content string
	timestamp time.Time
	isFromSender bool
}

type Model struct {
	Mode mode
	ViewportWidth int
	ViewportHeight int

	ctx context.Context
	db *db.DB
	theme theme.Theme
	input textinput.Model
	conversations []conversation
	conversationsIndex map[string]int
	selectedConversationIndex int
	messages map[string][]message
	conn *registry.Connection
}

func NewModel(ctx context.Context, db *db.DB, theme theme.Theme, conn *registry.Connection) *Model {
	dbConversations := db.ListConversations(ctx)
	conversations, conversationsIndex := mapConversations(dbConversations)

	messages := map[string][]message{}
	if len(conversations) > 0 {
		clientFingerprint := conversations[0].fingerprint
		dbMessages, _ := db.ListMessages(ctx, clientFingerprint)
		messages[clientFingerprint] = mapMessages(dbMessages)
	}

	ti := textinput.New()
	ti.Prompt = "> "
	ti.Placeholder = "type a message..."

	// Remove default styling for clean terminal look
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))

	return &Model{
		ctx: ctx,
		db: db,
		theme: theme,
		input: ti,
		conn: conn,
		conversations: conversations,
		conversationsIndex: conversationsIndex,
		messages: messages,
	}
}

type clientMessageEvent registry.MessageEvent

func readFromChannel(conn *registry.Connection) tea.Cmd {
	if conn == nil {
		return nil
	}

	return func() tea.Msg {
		return clientMessageEvent(conn.FetchMessage())
	}
}

func (m *Model) Init() tea.Cmd {
	m.Mode = Normal
	if len(m.conversations) > 0 {
		m.Mode = Insert
	}

	m.input.Focus()
	return tea.Batch(
		m.input.Cursor.BlinkCmd(),
		readFromChannel(m.conn),
	)
}

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.ViewportWidth = msg.Width
		m.ViewportHeight = msg.Height
	case clientMessageEvent:
		cMsgEvent := registry.MessageEvent(msg)
		m.updateConversations(cMsgEvent)
		m.updateChats(cMsgEvent, false)
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
		case "tab", "j", "down":
			m.moveToNextConversation()
		case "shift+tab", "k", "up":
			m.moveToPreviousConversation()
		case "enter":
			if text := m.input.Value(); text != "" {
				selectedConversation := m.conversations[m.selectedConversationIndex]
				messageEvent := registry.MessageEvent{
					Fingerprint: selectedConversation.fingerprint,
					Message: text,
				}

				m.input.SetValue("")

				m.updateChats(messageEvent, true)
				registry.HandleAdminMessage(messageEvent)

				createMessageRequest := db.CreateMessageRequest{
					ClientFingerprint: selectedConversation.fingerprint,
					SenderType: db.SenderAdmin,
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

func (m Model) View() string {
	if len(m.conversations) == 0 {
		prompt := "No messages. Press q to quit"
		return lipgloss.Place(m.ViewportWidth, m.ViewportHeight, lipgloss.Center, lipgloss.Center, prompt)
	}

	layout := lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.conversationsView(),
		m.chatPanelView(),
	)

	return lipgloss.Place(m.ViewportWidth, m.ViewportHeight, lipgloss.Center, lipgloss.Center, layout)
}

