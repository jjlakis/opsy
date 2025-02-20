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
//
// Agent Prompt (prompt.mdx):
//   - Contains the system prompt used by the AI agent
//   - Written in MDX format for rich text formatting
//   - Defines the agent's behavior, capabilities, and interaction patterns
//   - Can be customized per task via RunOptions.Prompt
//
// # Usage
//
// The assets are exposed through two embedded filesystems and one string variable:
//
//	var Themes embed.FS    // Access to theme configurations
//	var Tools embed.FS     // Access to tool configurations
//	var AgentPrompt string // Access to the agent's system prompt
//
// To access these assets in other parts of the application, import this package
// and use the appropriate embedded filesystem variable or string. The files can be
// read using standard fs.FS operations.
//
// Example:
//
//	themeData, err := assets.Themes.ReadFile("themes/default.yaml")
//	toolData, err := assets.Tools.ReadFile("tools/git.yaml")
//	prompt := assets.AgentPrompt // Direct access to the prompt string
package assets
