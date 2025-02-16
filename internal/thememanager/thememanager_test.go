package thememanager

import (
	"log/slog"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/assert"
)

// TestLoadTheme verifies theme loading functionality:
// - Loading default theme (empty name)
// - Loading theme by name
// - Handling invalid YAML format
// - Handling non-existent themes
func TestLoadTheme(t *testing.T) {
	tests := []struct {
		name    string
		theme   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "empty name (default theme)",
			theme:   "",
			wantErr: false,
		},
		{
			name:    "theme name only",
			theme:   "default",
			wantErr: false,
		},
		{
			name:    "non-existent theme",
			theme:   "nonexistent",
			wantErr: true,
			errMsg:  ErrThemeNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test logger
			testLogger := slog.New(slog.DiscardHandler)
			tm := New(WithLogger(testLogger))
			err := tm.LoadTheme(tt.theme)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
				return
			}

			assert.NoError(t, err)
			theme := tm.GetTheme()
			assert.NotNil(t, theme, "theme should not be nil")

			// Verify colors for valid themes
			colors := []struct {
				name  string
				color lipgloss.Color
			}{
				{"base.base00", theme.BaseColors.Base00},
				{"base.base01", theme.BaseColors.Base01},
				{"base.base02", theme.BaseColors.Base02},
				{"base.base03", theme.BaseColors.Base03},
				{"base.base04", theme.BaseColors.Base04},
				{"accent.accent0", theme.AccentColors.Accent0},
				{"accent.accent1", theme.AccentColors.Accent1},
				{"accent.accent2", theme.AccentColors.Accent2},
			}

			for _, c := range colors {
				assert.NotEmpty(t, string(c.color), "color %s should not be empty", c.name)
				if s := string(c.color); s != "" {
					assert.True(t, s[0] == '#', "color %s = %s, should start with #", c.name, s)
				}
			}
		})
	}
}

// TestThemeManager_WithLogger verifies logger functionality
func TestThemeManager_WithLogger(t *testing.T) {
	testLogger := slog.New(slog.DiscardHandler)
	tm := New(WithLogger(testLogger))

	assert.Equal(t, testLogger, tm.logger, "logger should be set correctly")
}

// TestThemeManager_WithDirectory verifies custom directory loading:
// - Loading themes from a custom directory
// - Handling non-existent directory
func TestThemeManager_WithDirectory(t *testing.T) {
	tests := []struct {
		name    string
		dir     string
		theme   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "custom directory",
			dir:     "testdata",
			theme:   "default",
			wantErr: true,
			errMsg:  ErrThemeNotFound,
		},
		{
			name:    "non-existent directory",
			dir:     "nonexistent",
			theme:   "default",
			wantErr: true,
			errMsg:  ErrThemeNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test logger
			testLogger := slog.New(slog.DiscardHandler)
			tm := New(
				WithDirectory(tt.dir),
				WithLogger(testLogger),
			)
			err := tm.LoadTheme(tt.theme)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, tm.GetTheme())
		})
	}
}
