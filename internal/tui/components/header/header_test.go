package header

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/datolabs-io/sredo/internal/thememanager"
	"github.com/stretchr/testify/assert"
)

func TestHeaderCreation(t *testing.T) {
	// Test default creation
	header := New()
	assert.NotNil(t, header)
	assert.Empty(t, header.task)

	// Test creation with options
	task := "Test Task"
	header = New(WithTask(task))
	assert.Equal(t, task, header.task)
}

func TestHeaderUpdate(t *testing.T) {
	header := New()

	// Test window size update
	newWidth := 100
	updatedHeader, _ := header.Update(tea.WindowSizeMsg{Width: newWidth})
	assert.Equal(t, newWidth, updatedHeader.maxWidth)
}

func TestHeaderView(t *testing.T) {
	theme := thememanager.Theme{
		BaseColors: thememanager.BaseColors{
			Base01: "#000000",
			Base04: "#FFFFFF",
		},
	}

	testCases := []struct {
		name     string
		task     string
		width    int
		contains []string
	}{
		{
			name:     "empty task",
			task:     "",
			width:    100,
			contains: []string{"Task:", ""},
		},
		{
			name:     "with task",
			task:     "Test Task",
			width:    100,
			contains: []string{"Task:", "Test Task"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			header := New(
				WithTask(tc.task),
				WithTheme(theme),
			)
			header.maxWidth = tc.width

			view := header.View()

			for _, expected := range tc.contains {
				assert.True(t,
					strings.Contains(view, expected),
					"Expected view to contain '%s'", expected,
				)
			}
		})
	}
}

func TestHeaderOptions(t *testing.T) {
	theme := thememanager.Theme{
		BaseColors: thememanager.BaseColors{
			Base01: "#000000",
			Base04: "#FFFFFF",
		},
	}

	testCases := []struct {
		name          string
		options       []Option
		expectedTask  string
		expectedTheme thememanager.Theme
	}{
		{
			name:          "with task option",
			options:       []Option{WithTask("Test Task")},
			expectedTask:  "Test Task",
			expectedTheme: thememanager.Theme{},
		},
		{
			name:          "with theme option",
			options:       []Option{WithTheme(theme)},
			expectedTask:  "",
			expectedTheme: theme,
		},
		{
			name: "with both options",
			options: []Option{
				WithTask("Test Task"),
				WithTheme(theme),
			},
			expectedTask:  "Test Task",
			expectedTheme: theme,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			header := New(tc.options...)
			assert.Equal(t, tc.expectedTask, header.task)
			assert.Equal(t, tc.expectedTheme, header.theme)
		})
	}
}
