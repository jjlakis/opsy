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

	//go:embed prompts/agent_system.mdx
	AgentSystemPrompt string
	//go:embed prompts/tool_system.mdx
	ToolSystemPrompt string
	//go:embed prompts/tool_user.mdx
	ToolUserPrompt string
)
