// Package footer provides a footer component for the terminal user interface.
//
// The footer component displays important information about the application's state,
// including:
//   - The AI engine being used (e.g., "Anthropic")
//   - Model configuration (model name, max tokens, temperature)
//   - Number of available tools
//   - Current status
//
// The component is built using the Bubble Tea framework and Lip Gloss styling
// library, providing a consistent look and feel with the rest of the application.
//
// Example usage:
//
//	footer := footer.New(
//	    footer.WithTheme(theme),
//	    footer.WithParameters(footer.Parameters{
//	        Engine:      "Anthropic",
//	        Model:       "claude-3-sonnet",
//	        MaxTokens:   1000,
//	        Temperature: 0.7,
//	        ToolsCount:  5,
//	    }),
//	)
package footer
