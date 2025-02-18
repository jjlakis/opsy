// Package toolmanager provides functionality for managing and executing tools
// within the application. It handles tool definitions, loading tools from files,
// and executing tool operations.
//
// Tools are defined using YAML configuration files and must implement the Tooling
// interface. Each tool has a name, display name, description, input schema, and
// execution logic.
//
// The package consists of two main components:
//
//   - Tool Manager: Handles loading and managing tool definitions from the filesystem
//   - Tool: Represents individual tools with their configurations and execution logic
//
// Example usage:
//
//	tm := toolmanager.New(
//		toolmanager.WithLogger(logger),
//		toolmanager.WithConfig(cfg),
//	)
//
//	if err := tm.LoadTools(); err != nil {
//		// Handle error
//	}
//
//	tools := tm.GetTools()
//
// Tool definitions are loaded from YAML files and include:
//   - Display name for UI presentation
//   - Description of the tool's functionality
//   - System prompt for AI interaction
//   - Input parameters with validation schemas
//
// All tools share common inputs including:
//   - task: The operation to perform
//   - working_directory: The directory context for the operation
//
// The package uses JSON Schema for input validation and provides error handling
// for common failure scenarios such as missing tools or invalid configurations.
package toolmanager
