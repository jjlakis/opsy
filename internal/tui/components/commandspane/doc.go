// Package commandspane provides a commands pane component for the terminal user interface.
//
// The commands pane component displays a scrollable list of executed commands, including:
//   - Timestamp of execution
//   - Working directory
//   - Command text
//
// Each command is styled to highlight:
//   - Working directory with a distinct background
//   - Command text in an accent color
//   - Timestamp in standard text color
//
// The component is built using the Bubble Tea framework and Lip Gloss styling
// library, providing a consistent look and feel with the rest of the application.
// It uses a viewport for scrollable content and supports dynamic resizing.
//
// Example usage:
//
//	commandspane := commandspane.New(
//	    commandspane.WithTheme(theme),
//	)
package commandspane
