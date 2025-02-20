package commandspane

import (
	"regexp"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/datolabs-io/sredo/internal/thememanager"
	"github.com/datolabs-io/sredo/internal/tool"
	"github.com/stretchr/testify/assert"
)

// stripANSI removes ANSI color codes from a string.
func stripANSI(str string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	return re.ReplaceAllString(str, "")
}

// TestNew tests the creation of a new commands pane component.
func TestNew(t *testing.T) {
	theme := thememanager.Theme{
		BaseColors: thememanager.BaseColors{
			Base01: "#000000",
			Base02: "#111111",
			Base03: "#222222",
			Base04: "#333333",
		},
		AccentColors: thememanager.AccentColors{
			Accent0: "#FF0000",
		},
	}

	m := New(
		WithTheme(theme),
	)

	assert.NotNil(t, m)
	assert.Equal(t, theme, m.theme)
	assert.NotNil(t, m.viewport)
	assert.Empty(t, m.commands)
}

// TestUpdate tests the update function of the commands pane component.
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

	// Test command message
	now := time.Now()
	testCmd := tool.Command{
		Command:          "ls -la",
		WorkingDirectory: "~/sredo",
		StartedAt:        now,
	}
	m, cmd = m.Update(testCmd)
	assert.Nil(t, cmd)
	assert.Len(t, m.commands, 1)
	assert.Equal(t, testCmd, m.commands[0])
}

// TestView tests the view function of the commands pane component.
func TestView(t *testing.T) {
	theme := thememanager.Theme{
		BaseColors: thememanager.BaseColors{
			Base01: "#000000",
			Base02: "#111111",
			Base03: "#222222",
			Base04: "#333333",
		},
		AccentColors: thememanager.AccentColors{
			Accent0: "#FF0000",
		},
	}

	m := New(
		WithTheme(theme),
	)

	// Set dimensions to test rendering
	m, _ = m.Update(tea.WindowSizeMsg{Width: 100, Height: 50})

	// Test initial view (empty commands)
	view := stripANSI(m.View())
	assert.NotEmpty(t, view)
	assert.Contains(t, view, "Commands")

	// Add test command
	now := time.Now()
	m.Update(tool.Command{
		Command:          "ls -la",
		WorkingDirectory: "~/sredo",
		StartedAt:        now,
	})

	// Test view with command
	view = stripANSI(m.View())
	assert.Contains(t, view, "Commands")
	assert.Contains(t, view, "~/sredo")
	assert.Contains(t, view, "ls -la")
	assert.Contains(t, view, now.Format("15:04:05"))
}

// TestInit tests the initialization of the commands pane component.
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
