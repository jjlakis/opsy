// Package assets provides embedded static assets for the sredo application.
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
//   - Tools configurations define how sredo interacts with various development tools
//   - Each tool can define its own system prompt in its configuration
//
// Prompts:
//
//   - AgentSystemPrompt (Main system prompt for the AI agent)
//     Used for task understanding and dispatching
//     Defines the agent's core behavior and capabilities
//     Controls how the agent interacts with tools and handles tasks
//
//   - ToolSystemPrompt (Common system prompt for all tools)
//     Appended to each tool's specific system prompt
//     Defines common behavior and patterns for all tools
//     Ensures consistent tool execution and output formatting
//
//   - ToolUserPrompt (User prompt template for tool execution)
//     Used to format commands for shell execution
//     Provides consistent command generation across tools
//     Includes task description and additional context
//
// # Usage
//
// The assets are exposed through two embedded filesystems and three string variables:
//
//	var Themes embed.FS         // Access to theme configurations
//	var Tools embed.FS          // Access to tool configurations
//	var AgentSystemPrompt string // Main system prompt for the AI agent
//	var ToolSystemPrompt string  // Common system prompt for all tools
//	var ToolUserPrompt string    // User prompt template for tool execution
//
// To access these assets in other parts of the application, import this package
// and use the appropriate embedded filesystem variable or string. The files can be
// read using standard fs.FS operations.
//
// Example:
//
//	themeData, err := assets.Themes.ReadFile("themes/default.yaml")
//	toolData, err := assets.Tools.ReadFile("tools/git.yaml")
//	agentPrompt := assets.AgentSystemPrompt
//	toolPrompt := assets.ToolSystemPrompt
package assets
