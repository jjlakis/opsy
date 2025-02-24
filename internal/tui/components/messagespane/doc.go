// Package messagespane provides a messages pane component for the terminal user interface.
//
// The messages pane component displays a scrollable list of messages, including:
//   - Agent messages (e.g., responses from the AI)
//   - Tool messages (e.g., output from executed commands)
//
// # Component Structure
//
// The Model type represents the messages pane component and provides the following methods:
//   - Init: Initializes the component (required by bubbletea.Model)
//   - Update: Handles messages and updates the component state
//   - View: Renders the component's current state
//   - WithTheme: Option function to set the theme for styling
//
// # Message Handling
//
// The component responds to:
//   - tea.WindowSizeMsg: Updates viewport dimensions and text wrapping
//   - agent.Message: Adds a new message to the pane
//
// Each message includes:
//   - Timestamp in [HH:MM:SS] format
//   - Source indicator ("Sredo" for agent, "Sredo->Tool" for tool messages)
//   - Message content with proper wrapping and formatting
//
// # Styling
//
// Each element is styled using dedicated styling methods:
//   - timestampStyle: formats the timestamp with a neutral color
//   - authorStyle: highlights the source (agent/tool) with distinct colors
//   - messageStyle: formats the message content with proper padding and background
//   - containerStyle: provides the overall pane styling with borders
//   - titleStyle: formats the "Messages" title
//
// Theme Integration:
//   - Base colors are used for backgrounds and text
//   - Accent colors differentiate between agent and tool messages
//   - All colors are configurable through the theme
//
// # Component Features
//
// The component automatically handles:
//   - Dynamic resizing of the viewport
//   - Message history accumulation
//   - Automatic scrolling to the latest message
//   - Proper text wrapping based on available width
//   - Different styling for agent vs tool messages
//   - Message sanitization (removing XML tags and extra whitespace)
//   - Viewport scrolling with mouse and keyboard
//
// # Thread Safety
//
// The messages pane component is safe for concurrent access:
//   - All updates are handled through message passing
//   - No internal mutable state is exposed
//   - Message list is only modified through the Update method
//   - Theme is immutable after creation
//
// # Viewport Controls
//
// The viewport supports standard scrolling controls:
//   - Mouse wheel: Scroll up/down
//   - PageUp/PageDown: Move by page
//   - Home/End: Jump to top/bottom
//   - Arrow keys: Scroll line by line
//
// Example usage:
//
//	// Create a new messages pane with theme
//	messagespane := messagespane.New(
//	    messagespane.WithTheme(theme),
//	)
//
//	// Initialize the component
//	cmd := messagespane.Init()
//
//	// Handle window resize
//	model, cmd := messagespane.Update(tea.WindowSizeMsg{Width: 100, Height: 50})
//
//	// Add a new message
//	model, cmd = messagespane.Update(agent.Message{
//	    Message:   "Hello, world!",
//	    Tool:      "",
//	    Timestamp: time.Now(),
//	})
//
//	// Add a tool message
//	model, cmd = messagespane.Update(agent.Message{
//	    Message:   "Running git status",
//	    Tool:      "Git",
//	    Timestamp: time.Now(),
//	})
//
//	// Render the component
//	view := messagespane.View()
package messagespane
