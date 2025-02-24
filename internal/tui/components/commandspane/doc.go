// Package commandspane provides a commands pane component for the terminal user interface.
//
// The commands pane component displays a scrollable list of executed commands, including:
//   - Timestamp of execution in [HH:MM:SS] format
//   - Working directory with a distinct background
//   - Command text in an accent color
//
// # Component Structure
//
// The Model type represents the commands pane component and provides the following methods:
//   - Init: Initializes the component (required by bubbletea.Model)
//   - Update: Handles messages and updates the component state
//   - View: Renders the component's current state
//
// The component supports configuration through options:
//   - WithTheme: Sets the theme for styling the component
//
// # Styling
//
// Each command is styled using dedicated styling methods:
//   - timestampStyle: formats the timestamp with a neutral color
//   - workdirStyle: highlights the working directory with a distinct background
//   - commandStyle: renders the command text in an accent color
//   - containerStyle: provides the overall pane styling with borders
//   - titleStyle: formats the "Commands" title
//
// Theme Integration:
//   - Base colors are used for backgrounds and borders
//   - Accent colors are used for command text highlighting
//   - All colors are configurable through the theme
//
// # Component Features
//
// The component automatically handles:
//   - Dynamic resizing of the viewport
//   - Command history accumulation
//   - Automatic scrolling to the latest command
//   - Proper text wrapping based on available width
//   - Long command wrapping with proper indentation
//   - Viewport scrolling for command history
//
// # Message Handling
//
// The component responds to:
//   - tea.WindowSizeMsg: Updates viewport dimensions
//   - tool.Command: Adds new command to history
//
// The component is built using the Bubble Tea framework and Lip Gloss styling
// library, providing a consistent look and feel with the rest of the application.
//
// Example usage:
//
//	commandspane := commandspane.New(
//	    commandspane.WithTheme(theme),
//	)
//
//	// Initialize the component
//	cmd := commandspane.Init()
//
//	// Handle window resize
//	model, cmd := commandspane.Update(tea.WindowSizeMsg{Width: 100, Height: 50})
//
//	// Add a new command
//	model, cmd = commandspane.Update(tool.Command{
//	    Command:          "ls -la",
//	    WorkingDirectory: "~/project",
//	    StartedAt:        time.Now(),
//	})
//
//	// Render the component
//	view := commandspane.View()
package commandspane
