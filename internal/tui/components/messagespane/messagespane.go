package messagespane

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/datolabs-io/sredo/internal/agent"
	"github.com/datolabs-io/sredo/internal/thememanager"
)

// Model represents the messages pane component.
type Model struct {
	theme          thememanager.Theme
	containerStyle lipgloss.Style
	textStyle      lipgloss.Style
	agentStyle     lipgloss.Style
	toolStyle      lipgloss.Style
	titleStyle     lipgloss.Style
	maxWidth       int
	maxHeight      int
	ready          bool
	viewport       viewport.Model
	messages       []string
}

// Option is a function that modifies the Model.
type Option func(*Model)

// New creates a new messages pane component.
func New(opts ...Option) *Model {
	m := &Model{
		viewport: viewport.New(0, 0),
	}

	for _, opt := range opts {
		opt(m)
	}

	m.containerStyle = containerStyle(m.theme)
	m.textStyle = textStyle(m.theme, m.maxWidth)
	m.agentStyle = agentStyle(m.theme)
	m.toolStyle = toolStyle(m.theme)
	m.titleStyle = titleStyle(m.theme)

	m.viewport.Style = m.textStyle

	return m
}

// title is the title of the messages pane.
const title = "Messages"

// Init initializes the messages pane component.
func (m *Model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the messages pane component.
func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.maxWidth = msg.Width
		m.maxHeight = msg.Height
		m.containerStyle = containerStyle(m.theme)
		m.textStyle = textStyle(m.theme, msg.Width)
		if !m.ready {
			m.ready = true
			m.viewport = viewport.New(msg.Width-6, msg.Height)
			m.viewport.SetContent(m.textStyle.Render(m.titleStyle.Render(title)))
		} else {
			m.viewport.Width = msg.Width - 6
			m.viewport.Height = m.maxHeight
		}
	case agent.Message:
		m.appendMessage(msg)
		m.viewport.GotoBottom()
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

// View renders the messages pane component.
func (m *Model) View() string {
	return m.containerStyle.Render(m.viewport.View())
}

// WithTheme sets the theme for the messages pane component.
func WithTheme(theme thememanager.Theme) Option {
	return func(m *Model) {
		m.theme = theme
	}
}

// containerStyle creates a style for the container of the messages pane component.
func containerStyle(theme thememanager.Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Background(theme.BaseColors.Base01).
		Padding(1, 2).
		Border(lipgloss.NormalBorder(), true).
		BorderForeground(theme.BaseColors.Base02).
		BorderBackground(theme.BaseColors.Base00).
		UnsetBorderBottom()
}

// textStyle creates a style for the text of the messages pane component.
func textStyle(theme thememanager.Theme, maxWidth int) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.BaseColors.Base03).
		Background(theme.BaseColors.Base01).
		Width(maxWidth - 6)
}

// agentStyle creates a style for agent messages.
func agentStyle(theme thememanager.Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.AccentColors.Accent1).
		Background(theme.BaseColors.Base01)
}

// toolStyle creates a style for tool messages.
func toolStyle(theme thememanager.Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.AccentColors.Accent2).
		Background(theme.BaseColors.Base01)
}

// titleStyle creates a style for the title.
func titleStyle(theme thememanager.Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.BaseColors.Base04).
		Background(theme.BaseColors.Base01).
		Bold(true)
}

// appendMessage formats and appends a message to the messages pane.
func (m *Model) appendMessage(msg agent.Message) {
	header := m.titleStyle.Render(title)
	timestamp := msg.Timestamp.Format("15:04:05")
	style := m.agentStyle
	author := agent.Name

	if msg.Tool != "" {
		author = fmt.Sprintf("%s->%s", agent.Name, msg.Tool)
		style = m.toolStyle
	}

	message := ":\n" + strings.ReplaceAll(msg.Message, "\n\n", "\n")
	message = fmt.Sprintf("%s%s", style.Bold(true).Render(author), style.Render(message))
	m.messages = append(m.messages, fmt.Sprintf("[%s] %s", timestamp, style.Render(message)))
	m.viewport.SetContent(m.textStyle.Render(fmt.Sprintf("%s\n\n%s", header, strings.Join(m.messages, "\n"))))
}
