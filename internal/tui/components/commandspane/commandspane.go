package commandspane

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/datolabs-io/sredo/internal/thememanager"
	"github.com/datolabs-io/sredo/internal/tool"
)

// Model represents the commands pane component.
type Model struct {
	theme     thememanager.Theme
	maxWidth  int
	maxHeight int
	viewport  viewport.Model
	commands  []tool.Command
}

// Option is a function that modifies the Model.
type Option func(*Model)

// title is the title of the commands pane.
const title = "Commands"

// New creates a new commands pane component.
func New(opts ...Option) *Model {
	m := &Model{
		viewport: viewport.New(0, 0),
		commands: []tool.Command{},
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

// Init initializes the commands pane component.
func (m *Model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the commands pane component.
func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.maxWidth = msg.Width - 6
		m.maxHeight = msg.Height
		m.viewport.Width = m.maxWidth
		m.viewport.Height = msg.Height
		m.viewport.Style = lipgloss.NewStyle().Background(m.theme.BaseColors.Base01)

		// Rerender all commands with new dimensions
		if len(m.commands) > 0 {
			m.renderCommands()
		} else {
			m.viewport.SetContent(m.titleStyle().Render(title))
		}
	case tool.Command:
		m.commands = append(m.commands, msg)
		m.renderCommands()
		m.viewport.GotoBottom()
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

// View renders the commands pane component.
func (m *Model) View() string {
	return m.containerStyle().Render(m.viewport.View())
}

// WithTheme sets the theme for the commands pane component.
func WithTheme(theme thememanager.Theme) Option {
	return func(m *Model) {
		m.theme = theme
	}
}

// containerStyle creates a style for the container of the commands pane component.
func (m *Model) containerStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(m.theme.BaseColors.Base01).
		Padding(1, 2).
		Border(lipgloss.NormalBorder(), true).
		BorderForeground(m.theme.BaseColors.Base02).
		BorderBackground(m.theme.BaseColors.Base00)
}

// commandStyle creates a style for the command text.
func (m *Model) commandStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(m.theme.AccentColors.Accent0).
		Background(m.theme.BaseColors.Base01)
}

// timestampStyle creates a style for the timestamp of the commands pane component.
func (m *Model) timestampStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(m.theme.BaseColors.Base03).
		Background(m.theme.BaseColors.Base01).
		PaddingRight(1)
}

// workdirStyle creates a style for the working directory.
func (m *Model) workdirStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(m.theme.BaseColors.Base04).
		Background(m.theme.BaseColors.Base03).
		Margin(0, 1, 0, 0).
		MarginBackground(m.theme.BaseColors.Base01).
		Padding(0, 1)
}

// titleStyle creates a style for the title.
func (m *Model) titleStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(m.theme.BaseColors.Base04).
		Background(m.theme.BaseColors.Base01).
		Bold(true).
		Width(m.maxWidth)
}

// renderCommands formats and renders all commands
func (m *Model) renderCommands() {
	output := strings.Builder{}
	output.WriteString(m.titleStyle().Render(title))
	output.WriteString("\n\n")

	for _, cmd := range m.commands {
		timestamp := m.timestampStyle().Render(fmt.Sprintf("[%s]", cmd.StartedAt.Format("15:04:05")))
		workdir := m.workdirStyle().Render(cmd.WorkingDirectory)

		commandStyle := m.commandStyle().Width(m.maxWidth - lipgloss.Width(timestamp) - lipgloss.Width(workdir))
		command := commandStyle.Render(cmd.Command)

		output.WriteString(fmt.Sprintf("%s%s%s", timestamp, workdir, command))
		output.WriteString("\n\n")
	}

	m.viewport.SetContent(output.String())
}
