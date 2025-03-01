// Package assets provides embedded static assets for the opsy application.
//
// The package uses Go's embed functionality to include various static assets
// that are required for the application to function. These assets are compiled
// into the binary, ensuring they are always available at runtime.
//
// # Embedded Assets
//
// The package contains three main categories of embedded assets:
//
// Themes Directory (/themes):
//   - Contains theme configuration files in YAML format
//   - Includes default.yaml which defines the default application theme
//   - Themes are used to customize the appearance of the terminal UI
//
// Tools Directory (/tools):
//   - Contains tool-specific configuration files in YAML format
//   - Includes git.yaml which defines Git-related configurations and commands
//   - Tools configurations define how opsy interacts with various development tools
//   - Each tool can define its own system prompt in its configuration
//
// Prompts:
//
//   - Agent System Prompt (Main system prompt for the AI agent)
//     Used for task understanding and dispatching
//     Defines the agent's core behavior and capabilities
//     Controls how the agent interacts with tools and handles tasks
//
//   - Tool System Prompt (Common system prompt for all tools)
//     Appended to each tool's specific system prompt
//     Defines common behavior and patterns for all tools
//     Ensures consistent tool execution and output formatting
//
//   - Tool User Prompt (User prompt template for tool execution)
//     Used to format commands for shell execution
//     Provides consistent command generation across tools
//     Includes task description and additional context
//
// # Usage
//
// The assets are exposed through two embedded filesystems and prompt rendering functions:
//
//	var Themes embed.FS  // Access to theme configurations
//	var Tools embed.FS   // Access to tool configurations
//
// To access theme and tool configurations, use standard fs.FS operations:
//
//	themeData, err := assets.Themes.ReadFile("themes/default.yaml")
//	toolData, err := assets.Tools.ReadFile("tools/git.yaml")
//
// To render prompts, use the provided render functions:
//
//	// Render agent system prompt
//	prompt, err := assets.RenderAgentSystemPrompt(&AgentSystemPromptData{
//		Shell: "/bin/bash",
//	})
//
//	// Render tool system prompt
//	prompt, err := assets.RenderToolSystemPrompt(&ToolSystemPromptData{
//		Shell:      "/bin/bash",
//		Name:       "git",
//		Executable: "/usr/bin/git",
//		Rules:      []string{"rule1", "rule2"},
//	})
//
//	// Render tool user prompt
//	prompt, err := assets.RenderToolUserPrompt(&ToolUserPromptData{
//		Task:             "Clone repository",
//		Params:          map[string]any{"url": "https://github.com/example/repo"},
//		Context:         map[string]string{"branch": "main"},
//		WorkingDirectory: "/path/to/workspace",
//	})
//
// Each render function accepts a specific data struct and returns the rendered prompt
// as a string. If there's an error during rendering, it will be returned along with
// an empty string.
package assets
