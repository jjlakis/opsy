// Package commandspane provides a commands pane component for the terminal user interface.
//
// The commands pane component displays a scrollable list of executed commands, including:
//   - Timestamp of execution in [HH:MM:SS] format
//   - Working directory with a distinct background
//   - Command text in an accent color
//
// Each command is styled using dedicated styling methods:
//   - timestampStyle: formats the timestamp with a neutral color
//   - workdirStyle: highlights the working directory with a distinct background
//   - commandStyle: renders the command text in an accent color
//   - containerStyle: provides the overall pane styling with borders
//   - titleStyle: formats the "Commands" title
//
// The component automatically handles:
//   - Dynamic resizing of the viewport
//   - Command history accumulation
//   - Automatic scrolling to the latest command
//   - Proper text wrapping based on available width
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
