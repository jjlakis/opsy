// Package messagespane provides a messages pane component for the terminal user interface.
//
// The messages pane component displays a scrollable list of messages, including:
//   - Agent messages (e.g., responses from the AI)
//   - Tool messages (e.g., output from executed commands)
//
// Each message includes:
//   - Timestamp
//   - Source (Agent or Tool)
//   - Content
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
