package commandspane

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/datolabs-io/sredo/internal/thememanager"
)

// Model represents the commands pane component.
type Model struct {
	theme          thememanager.Theme
	containerStyle lipgloss.Style
	textStyle      lipgloss.Style
	workdirStyle   lipgloss.Style
	commandStyle   lipgloss.Style
	titleStyle     lipgloss.Style
	maxWidth       int
	maxHeight      int
	viewport       viewport.Model
}

// Option is a function that modifies the Model.
type Option func(*Model)

// New creates a new commands pane component.
func New(opts ...Option) *Model {
	m := &Model{
		viewport: viewport.New(0, 0),
	}

	for _, opt := range opts {
		opt(m)
	}

	m.containerStyle = containerStyle(m.theme)
	m.textStyle = textStyle(m.theme, m.maxWidth)
	m.workdirStyle = workdirStyle(m.theme)
	m.commandStyle = commandStyle(m.theme)
	m.titleStyle = titleStyle(m.theme)

	m.viewport.Style = m.textStyle

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
		m.maxWidth = msg.Width
		m.maxHeight = msg.Height
		m.containerStyle = containerStyle(m.theme)
		m.textStyle = textStyle(m.theme, msg.Width)
		m.viewport.Width = msg.Width - 6
		m.viewport.Height = m.maxHeight
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

// View renders the commands pane component.
func (m *Model) View() string {
	content := m.titleStyle.Render("Commands") + "\n\n" + m.commandView()
	m.viewport.SetContent(m.textStyle.Render(content))

	return m.containerStyle.Render(m.viewport.View())
}

// commandView renders a command with its working directory.
func (m *Model) commandView() string {
	timestamp := fmt.Sprintf("[%s] ", time.Now().Format("15:04:05"))
	workingDirectory := m.workdirStyle.Render("~/sredo")
	command := m.commandStyle.Render("ls -la")

	return fmt.Sprintf("%s%s%s", timestamp, workingDirectory, command)
}

// WithTheme sets the theme for the commands pane component.
func WithTheme(theme thememanager.Theme) Option {
	return func(m *Model) {
		m.theme = theme
	}
}

// containerStyle creates a style for the container of the commands pane component.
func containerStyle(theme thememanager.Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Background(theme.BaseColors.Base01).
		Padding(1, 2).
		Border(lipgloss.NormalBorder(), true).
		BorderForeground(theme.BaseColors.Base02).
		BorderBackground(theme.BaseColors.Base00)
}

// textStyle creates a style for the text of the commands pane component.
func textStyle(theme thememanager.Theme, maxWidth int) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.BaseColors.Base03).
		Background(theme.BaseColors.Base01).
		Width(maxWidth - 6)
}

// workdirStyle creates a style for the working directory.
func workdirStyle(theme thememanager.Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.BaseColors.Base04).
		Background(theme.BaseColors.Base03).
		Margin(0, 1, 0, 0).
		MarginBackground(theme.BaseColors.Base01).
		Padding(0, 1)
}

// commandStyle creates a style for the command text.
func commandStyle(theme thememanager.Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.AccentColors.Accent0).
		Background(theme.BaseColors.Base01)
}

// titleStyle creates a style for the title.
func titleStyle(theme thememanager.Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.BaseColors.Base04).
		Background(theme.BaseColors.Base01)
}
