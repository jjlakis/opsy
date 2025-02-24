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
//   - Managing the exec tool as a special built-in tool
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
// Tool Validation:
//
// Each tool definition is validated to ensure:
//   - Display name is provided and non-empty
//   - Description is provided and non-empty
//   - Input parameters have valid types and descriptions
//   - System prompt is valid if provided
//   - Executable path exists and is executable if specified
//
// Exec Tool:
//
// The exec tool is a special built-in tool that:
//   - Is always loaded regardless of configuration
//   - Provides direct command execution capabilities
//   - Uses the shell specified in configuration
//   - Has its own timeout configuration
//
// Error Handling:
//
// The package uses the following error constants:
//   - ErrLoadingTools: Returned when tools cannot be loaded from directory
//   - ErrLoadingTool: Returned when a specific tool fails to load
//   - ErrParsingTool: Returned when tool YAML parsing fails
//   - ErrToolNotFound: Returned when requested tool doesn't exist
//   - ErrInvalidToolDefinition: Returned when tool definition is invalid
//
// Thread Safety:
//
// The toolmanager is safe for concurrent access:
//   - Tool loading is synchronized
//   - Tool access methods are safe for concurrent use
//   - Tool instances are immutable after creation
//   - The exec tool maintains its own thread safety
//
// The toolmanager requires an agent to be provided for tool execution. The agent
// is responsible for running tool operations and managing their lifecycle.
//
// The package uses JSON Schema for input validation and provides error handling
// for common failure scenarios such as missing tools or invalid configurations.
package toolmanager
