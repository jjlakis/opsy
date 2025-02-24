// Package tui provides the terminal user interface for the Sredo application.
//
// The TUI is built using the Bubble Tea framework and consists of four main components:
//   - Header: Displays the current task and application state
//   - Messages Pane: Shows the conversation between the user and the AI
//   - Commands Pane: Displays executed commands and their output
//   - Footer: Shows AI model configuration and status
//
// Each component is independently managed and styled, using the application's theme
// for consistent appearance. The layout automatically adjusts to the terminal size,
// with dynamic height calculations:
//   - Header height adjusts based on task text wrapping
//   - Messages pane takes 2/3 of the remaining height
//   - Commands pane takes 1/3 of the remaining height
//   - Footer maintains a fixed height
//
// Example usage:
//
//	tui := tui.New(
//	    tui.WithTheme(theme),
//	    tui.WithConfig(cfg),
//	    tui.WithTask("Analyze system performance"),
//	    tui.WithToolsCount(5),
//	)
//	p := tea.NewProgram(tui)
//	if _, err := p.Run(); err != nil {
//	    log.Fatal(err)
//	}
//
// The TUI can be configured using functional options:
//   - WithTheme: Sets the theme for all components
//   - WithConfig: Sets the AI model configuration
//   - WithTask: Sets the current task being executed
//   - WithToolsCount: Sets the number of available tools
//
// Message Handling:
//
// The TUI processes several types of messages:
//   - tea.WindowSizeMsg: Triggers layout recalculation
//   - tea.KeyMsg: Handles keyboard input (e.g., Ctrl+C for quit)
//   - agent.Message: Updates the messages pane
//   - tool.Command: Updates the commands pane
//   - agent.Status: Updates the footer status
//
// Thread Safety:
//
// The TUI is designed to be thread-safe:
//   - All updates are handled through the message system
//   - Components maintain their own state
//   - No shared mutable state between components
//   - Safe for concurrent message processing
//
// All components receive and handle window size messages to maintain proper layout,
// and each component can process its own specific messages for additional functionality.
package tui
