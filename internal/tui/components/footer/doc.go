// Package footer provides a footer component for the terminal user interface.
//
// The footer component displays important information about the application's state,
// including:
//   - The AI engine being used (e.g., "Anthropic")
//   - Model configuration (model name, max tokens, temperature)
//   - Number of available tools
//   - Current status
//
// # Component Structure
//
// The Model type represents the footer component and provides the following methods:
//   - Init: Initializes the component (required by bubbletea.Model)
//   - Update: Handles messages and updates the component state
//   - View: Renders the component's current state
//
// The component supports configuration through options:
//   - WithTheme: Sets the theme for styling the component
//   - WithParameters: Sets the application parameters to display
//
// # Message Handling
//
// The component responds to:
//   - tea.WindowSizeMsg: Updates viewport dimensions
//   - agent.Status: Updates the current status display
//
// # Styling
//
// Each element is styled using dedicated styling methods:
//   - containerStyle: provides the overall footer styling with background
//   - textStyle: formats the text content with appropriate colors
//
// Theme Integration:
//   - Base colors are used for backgrounds and text
//   - All colors are configurable through the theme
//
// # Component Features
//
// The component automatically handles:
//   - Dynamic resizing based on window width
//   - Status updates through message passing
//   - Right-aligned status display
//   - Bold labels with regular value text
//   - Proper spacing and padding
//
// # Thread Safety
//
// The footer component is safe for concurrent access:
//   - All updates are handled through message passing
//   - No internal mutable state is exposed
//   - Theme and parameters are immutable after creation
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
//
//	// Initialize the component
//	cmd := footer.Init()
//
//	// Handle window resize
//	model, cmd := footer.Update(tea.WindowSizeMsg{Width: 100})
//
//	// Update status
//	model, cmd = footer.Update(agent.Status("Running"))
//
//	// Render the component
//	view := footer.View()
package footer
