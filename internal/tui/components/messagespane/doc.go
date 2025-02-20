// Package messagespane provides a messages pane component for the terminal user interface.
//
// The messages pane component displays a scrollable list of messages, including:
//   - Agent messages (e.g., responses from the AI)
//   - Tool messages (e.g., output from executed commands)
//
// Each message includes:
//   - Timestamp in [HH:MM:SS] format
//   - Source indicator ("Sredo" for agent, "Sredo->Tool" for tool messages)
//   - Message content with proper wrapping and formatting
//
// Each message is styled using dedicated styling methods:
//   - timestampStyle: formats the timestamp with a neutral color
//   - authorStyle: highlights the source (agent/tool) with distinct colors
//   - messageStyle: formats the message content with proper padding and background
//   - containerStyle: provides the overall pane styling with borders
//   - titleStyle: formats the "Messages" title
//
// The component automatically handles:
//   - Dynamic resizing of the viewport
//   - Message history accumulation
//   - Automatic scrolling to the latest message
//   - Proper text wrapping based on available width
//   - Different styling for agent vs tool messages
//
// The component is built using the Bubble Tea framework and Lip Gloss styling
// library, providing a consistent look and feel with the rest of the application.
// It uses a viewport for scrollable content and supports dynamic resizing.
//
// Example usage:
//
//	messagespane := messagespane.New(
//	    messagespane.WithTheme(theme),
//	)
package messagespane
