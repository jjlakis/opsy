package footer

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/datolabs-io/sredo/internal/thememanager"
	"github.com/stretchr/testify/assert"
)

// TestNew tests the creation of a new footer component.
func TestNew(t *testing.T) {
	theme := thememanager.Theme{
		BaseColors: thememanager.BaseColors{
			Base01: "#000000",
			Base04: "#FFFFFF",
		},
	}
	params := Parameters{
		Engine:      "TestEngine",
		Model:       "TestModel",
		MaxTokens:   1000,
		Temperature: 0.7,
		ToolsCount:  5,
	}

	m := New(
		WithTheme(theme),
		WithParameters(params),
	)

	assert.NotNil(t, m)
	assert.Equal(t, params, m.parameters)
	assert.Equal(t, theme, m.theme)
	assert.Equal(t, "Ready", m.status)
}

// TestUpdate tests the update function of the footer component.
func TestUpdate(t *testing.T) {
	theme := thememanager.Theme{
		BaseColors: thememanager.BaseColors{
			Base01: "#000000",
			Base04: "#FFFFFF",
		},
	}
	m := New(WithTheme(theme))

	// Test window size message
	newModel, cmd := m.Update(tea.WindowSizeMsg{Width: 100})
	assert.NotNil(t, newModel)
	assert.Nil(t, cmd)
	assert.Equal(t, 100, newModel.maxWidth)
}

// TestView tests the view function of the footer component.
func TestView(t *testing.T) {
	theme := thememanager.Theme{
		BaseColors: thememanager.BaseColors{
			Base01: "#000000",
			Base04: "#FFFFFF",
		},
	}
	params := Parameters{
		Engine:      "TestEngine",
		Model:       "TestModel",
		MaxTokens:   1000,
		Temperature: 0.7,
		ToolsCount:  5,
	}

	m := New(
		WithTheme(theme),
		WithParameters(params),
	)
	m.maxWidth = 100

	view := m.View()
	assert.NotEmpty(t, view)
	assert.Contains(t, view, "TestEngine")
	assert.Contains(t, view, "TestModel")
	assert.Contains(t, view, "Ready") // Default status
}

// TestInit tests the initialization of the footer component.
func TestInit(t *testing.T) {
	theme := thememanager.Theme{
		BaseColors: thememanager.BaseColors{
			Base01: "#000000",
			Base04: "#FFFFFF",
		},
	}
	m := New(WithTheme(theme))
	cmd := m.Init()
	assert.Nil(t, cmd)
}
