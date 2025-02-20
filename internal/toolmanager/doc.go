// Package toolmanager provides functionality for managing tools within the application.
// It handles loading tool definitions from files and managing their lifecycle.
//
// Tools are defined using YAML configuration files and must implement the tool.Tool
// interface from the tool package. The toolmanager loads these definitions and
// creates the appropriate tool instances.
//
// The toolmanager is responsible for:
//   - Loading tool definitions from YAML files
//   - Creating and managing tool instances
//   - Providing access to tools by name
//   - Maintaining the tool registry
//
// Example usage:
//
//	agent := agent.New(
//		agent.WithLogger(logger),
//		agent.WithConfig(cfg),
//	)
//
//	tm := toolmanager.New(
//		toolmanager.WithLogger(logger),
//		toolmanager.WithConfig(cfg),
//		toolmanager.WithAgent(agent),
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
//   - Optional executable path for command-line tools
//
// The toolmanager supports loading both regular tools (defined in YAML) and
// built-in tools like the exec tool. Each tool must conform to the tool.Tool
// interface and is validated during loading.
//
// The toolmanager requires an agent to be provided for tool execution. The agent
// is responsible for running tool operations and managing their lifecycle.
//
// The package uses JSON Schema for input validation and provides error handling
// for common failure scenarios such as missing tools or invalid configurations.
package toolmanager
