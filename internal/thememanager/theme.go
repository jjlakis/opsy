package thememanager

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v3"
)

const (
	// ErrMissingColors is the error message for missing required colors.
	ErrMissingColors = "missing required colors"
	// ErrDecodingTheme is returned when theme decoding fails.
	ErrDecodingTheme = "failed to decode theme"
)

// Theme defines the color palette for the application TUI.
type Theme struct {
	// BaseColors contains the base color palette.
	BaseColors BaseColors `yaml:"base"`
	// AccentColors contains the accent color palette.
	AccentColors AccentColors `yaml:"accent"`
}

// BaseColors contains the base color palette.
type BaseColors struct {
	// Base00 is used for primary background.
	Base00 lipgloss.Color `yaml:"base00"`
	// Base01 is used for secondary background (status bars, input).
	Base01 lipgloss.Color `yaml:"base01"`
	// Base02 is used for borders and dividers.
	Base02 lipgloss.Color `yaml:"base02"`
	// Base03 is used for muted or disabled text.
	Base03 lipgloss.Color `yaml:"base03"`
	// Base04 is used for primary text content.
	Base04 lipgloss.Color `yaml:"base04"`
}

// AccentColors contains the accent color palette.
type AccentColors struct {
	// Accent0 is used for command text and prompts.
	Accent0 lipgloss.Color `yaml:"accent0"`
	// Accent1 is used for agent messages and success states.
	Accent1 lipgloss.Color `yaml:"accent1"`
	// Accent2 is used for tool output and links.
	Accent2 lipgloss.Color `yaml:"accent2"`
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (t *Theme) UnmarshalYAML(value *yaml.Node) error {
	type ThemeYAML Theme
	var tmp ThemeYAML

	if err := value.Decode(&tmp); err != nil {
		// Wrap the YAML error with our error message
		return fmt.Errorf("%s: %v", ErrDecodingTheme, err)
	}

	required := []struct {
		name  string
		color lipgloss.Color
	}{
		{"base.base00", tmp.BaseColors.Base00},
		{"base.base01", tmp.BaseColors.Base01},
		{"base.base02", tmp.BaseColors.Base02},
		{"base.base03", tmp.BaseColors.Base03},
		{"base.base04", tmp.BaseColors.Base04},
		{"accent.accent0", tmp.AccentColors.Accent0},
		{"accent.accent1", tmp.AccentColors.Accent1},
		{"accent.accent2", tmp.AccentColors.Accent2},
	}

	var missing []string
	for _, r := range required {
		if r.color == "" {
			missing = append(missing, r.name)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("%s: %v", ErrMissingColors, missing)
	}

	*t = Theme(tmp)
	return nil
}
