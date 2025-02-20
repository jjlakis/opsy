package messagespane

import (
	"regexp"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/datolabs-io/sredo/internal/agent"
	"github.com/datolabs-io/sredo/internal/thememanager"
	"github.com/stretchr/testify/assert"
)

// stripANSI removes ANSI color codes from a string.
func stripANSI(str string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	return re.ReplaceAllString(str, "")
}

// TestNew tests the creation of a new messages pane component.
func TestNew(t *testing.T) {
	theme := thememanager.Theme{
		BaseColors: thememanager.BaseColors{
			Base01: "#000000",
			Base02: "#111111",
			Base03: "#222222",
			Base04: "#333333",
		},
		AccentColors: thememanager.AccentColors{
			Accent1: "#FF0000",
			Accent2: "#00FF00",
		},
	}

	m := New(
		WithTheme(theme),
	)

	assert.NotNil(t, m)
	assert.Equal(t, theme, m.theme)
	assert.NotNil(t, m.viewport)
	assert.Empty(t, m.messages)
}

// TestUpdate tests the update function of the messages pane component.
func TestUpdate(t *testing.T) {
	theme := thememanager.Theme{
		BaseColors: thememanager.BaseColors{
			Base01: "#000000",
			Base02: "#111111",
			Base03: "#222222",
			Base04: "#333333",
		},
	}
	m := New(WithTheme(theme))

	// Test window size message
	newModel, cmd := m.Update(tea.WindowSizeMsg{Width: 100, Height: 50})
	assert.NotNil(t, newModel)
	assert.Nil(t, cmd)
	assert.Equal(t, 94, newModel.maxWidth) // Width - 6 for padding
	assert.Equal(t, 50, newModel.maxHeight)
	assert.Equal(t, 94, newModel.viewport.Width)
	assert.Equal(t, 50, newModel.viewport.Height)

	// Test message handling
	testMsg := agent.Message{
		Message:   "Test message",
		Tool:      "",
		Timestamp: time.Now(),
	}
	m, cmd = m.Update(testMsg)
	assert.Nil(t, cmd)
	assert.Len(t, m.messages, 1)
	assert.Equal(t, testMsg, m.messages[0])
}

// TestView tests the view function of the messages pane component.
func TestView(t *testing.T) {
	theme := thememanager.Theme{
		BaseColors: thememanager.BaseColors{
			Base01: "#000000",
			Base02: "#111111",
			Base03: "#222222",
			Base04: "#333333",
		},
		AccentColors: thememanager.AccentColors{
			Accent1: "#FF0000",
			Accent2: "#00FF00",
		},
	}

	m := New(
		WithTheme(theme),
	)

	// Set dimensions to test rendering
	m, _ = m.Update(tea.WindowSizeMsg{Width: 100, Height: 50})

	// Test initial view (empty messages)
	view := stripANSI(m.View())
	assert.NotEmpty(t, view)
	assert.Contains(t, view, "Messages")

	// Add test messages
	now := time.Now()
	m.Update(agent.Message{
		Message:   "Hello",
		Tool:      "",
		Timestamp: now,
	})
	m.Update(agent.Message{
		Message:   "Running git command",
		Tool:      "Git",
		Timestamp: now,
	})

	// Test view with messages
	view = stripANSI(m.View())
	assert.Contains(t, view, "Messages")
	assert.Contains(t, view, "Sredo:")
	assert.Contains(t, view, "Sredo->Git:")
	assert.Contains(t, view, "Hello")
	assert.Contains(t, view, "Running git command")
}

// TestInit tests the initialization of the messages pane component.
func TestInit(t *testing.T) {
	theme := thememanager.Theme{
		BaseColors: thememanager.BaseColors{
			Base01: "#000000",
		},
	}
	m := New(WithTheme(theme))
	cmd := m.Init()
	assert.Nil(t, cmd)
}
