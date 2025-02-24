/*
Package tool provides functionality for defining and executing tools within the sredo application.

A tool is a unit of functionality that can be executed by the agent to perform specific tasks.
Each tool has a definition that describes its capabilities, inputs, and behavior.

# Tool Definition

Tools are defined using the Definition struct, which includes:

  - DisplayName: Human-readable name shown in the UI
  - Description: Detailed description of the tool's purpose
  - Rules: Additional rules the tool must follow
  - Inputs: Map of input parameters the tool accepts
  - Executable: Optional path to an executable the tool uses

# Input Schema

Tools can define their input requirements using the Input struct:

  - Type: Data type of the input (e.g., "string", "number")
  - Description: Human-readable description of the input
  - Default: Default value if none is provided
  - Examples: List of example values
  - Optional: Whether the input is required

Every tool automatically includes common inputs:

  - task: The task to be executed (required)
  - working_directory: Directory to execute in (optional, defaults to ".")
  - context: Additional context parameters (optional)

# Tool Interface

The Tool interface defines the methods a tool must implement:

  - GetName: Returns the tool's identifier
  - GetDisplayName: Returns the human-readable name
  - GetDescription: Returns the tool's description
  - GetInputSchema: Returns the JSON schema for inputs
  - Execute: Executes the tool with given inputs

# Tool Types

The package includes two main types of tools:

1. Regular tools (tool): Base implementation that can be extended
2. Exec tools (execTool): Special tools that execute shell commands

The exec tool has specific features:

  - Command execution with configurable timeouts
  - Working directory resolution (absolute, relative, and ./ paths)
  - Command output and exit code capture
  - Timestamp tracking for command execution
  - Process group management for proper cleanup

# Example Usage

Creating a new tool:

	def := tool.Definition{
		DisplayName: "My Tool",
		Description: "Does something useful",
		Rules: []string{"Follow these rules"},
		Inputs: map[string]tool.Input{
			"param": {
				Type: "string",
				Description: "A parameter",
				Optional: false,
			},
		},
	}

	myTool := tool.New("my-tool", def, logger, cfg, runner)

Creating an exec tool:

	execTool := tool.NewExecTool(logger, cfg)

Using the exec tool:

	output, err := execTool.Execute(map[string]any{
		"command": "ls -la",
		"working_directory": "./mydir",
	}, ctx)

# Error Handling

The package defines several error types for validation:

  - ErrToolMissingDisplayName: Tool definition lacks a display name
  - ErrToolMissingDescription: Tool definition lacks a description
  - ErrToolInputMissingType: Input definition lacks a type
  - ErrToolInputMissingDescription: Input definition lacks a description
  - ErrToolExecutableNotFound: Specified executable not found
  - ErrInvalidToolInputType: Input value has wrong type

# Thread Safety

Tools are designed to be thread-safe and can be executed concurrently.
Each execution:
  - Gets its own context and timeout based on configuration
  - Has isolated working directory resolution
  - Maintains independent command state and output
  - Uses process groups for clean termination
*/
package tool
