package assets

import (
	"embed"
)

var (
	//go:embed themes
	Themes embed.FS
	//go:embed tools
	Tools embed.FS
	// ToolsDir is the directory containing the tools.
	ToolsDir = "tools"
	// ThemeDir is the directory containing the themes.
	ThemeDir = "themes"
)
