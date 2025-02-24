// Package header provides a header component for the terminal user interface.
// The header displays the current task and can be styled using themes.
//
// # Component Structure
//
// The Model type represents the header component and provides the following methods:
//   - Init: Initializes the component (required by bubbletea.Model)
//   - Update: Handles messages and updates the component state
//   - View: Renders the component's current state
//
// The component supports configuration through options:
//   - WithTask: Sets the task text to display
//   - WithTheme: Sets the theme for styling the component
//
// # Message Handling
//
// The component responds to:
//   - tea.WindowSizeMsg: Updates viewport dimensions and text wrapping
//
// # Styling
//
// Each element is styled using dedicated styling methods:
//   - containerStyle: provides the overall header styling with background
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
//   - Text wrapping for long task descriptions
//   - Bold label with regular task text
//   - Proper spacing and padding
//   - Theme-based styling
//
// # Thread Safety
//
// The header component is safe for concurrent access:
//   - All updates are handled through message passing
//   - No internal mutable state is exposed
//   - Theme and task text are immutable after creation
//
// Example usage:
//
//	header := header.New(
//		header.WithTask("Current Task"),
//		header.WithTheme(myTheme),
//	)
//
//	// Initialize the component
//	cmd := header.Init()
//
//	// Handle window resize
//	model, cmd := header.Update(tea.WindowSizeMsg{Width: 100})
//
//	// Render the component
//	view := header.View()
package header
