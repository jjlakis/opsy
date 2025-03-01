/*
Package agent provides functionality for executing tasks using AI-powered tools within the opsy application.

The agent acts as a bridge between the user's task requests and the available tools, using the Anthropic
Claude API to intelligently select and execute appropriate tools based on the task requirements.

# Core Components

The package consists of several key components:

  - Agent: The main struct that handles task execution and tool management
  - Communication: Channels for sending messages, commands, and status updates
  - Message: Represents a message from the agent or tool execution
  - Status: Represents the current state of the agent (Running, Finished, etc.)

# Agent Configuration

The agent can be configured using functional options:

	agent := agent.New(
		agent.WithConfig(cfg),
		agent.WithLogger(logger),
		agent.WithContext(ctx),
		agent.WithCommunication(comm),
	)

Available options include:
  - WithConfig: Sets the configuration for the agent
  - WithLogger: Sets the logger for the agent
  - WithContext: Sets the context for the agent
  - WithCommunication: Sets the communication channels

# Task Execution

Tasks are executed using the Run method:

	outputs, err := agent.Run(&tool.RunOptions{
		Task:   "Clone the repository",
		Tools:  toolManager.GetTools(),
		Prompt: customPrompt, // Optional: Override default system prompt
		Caller: "git",       // Optional: Tool identifier for messages
	}, ctx)

The agent will:
1. Parse the task and available tools
2. Use the Anthropic API to determine which tools to use
3. Execute the selected tools with appropriate parameters
4. Return the combined output from all tool executions

The agent supports customizing the system prompt through RunOptions.Prompt,
which allows overriding the default behavior when needed.

# Communication

The agent uses channels to communicate its progress:

  - Messages: Task progress and tool output messages
  - Commands: Commands executed by tools
  - Status: Current agent status (Running, Finished)

Example usage:

	comm := &agent.Communication{
		Commands: make(chan tool.Command),
		Messages: make(chan agent.Message),
		Status:   make(chan agent.Status),
	}

	go func() {
		for msg := range comm.Messages {
			// Handle message
		}
	}()

# Tool Integration

Tools are converted to a format compatible with the Anthropic API:

  - Name: Tool identifier
  - Description: Tool purpose and functionality
  - InputSchema: JSON Schema defining valid inputs

The agent ensures proper conversion and validation of tools before use.
By default, parallel tool use is disabled to ensure deterministic execution.

# Error Handling

The package defines several error types:

  - ErrNoRunOptions: No options provided for Run
  - ErrNoTaskProvided: No task specified in options

All errors are properly logged with contextual information using structured logging.
Tool execution errors are captured and reflected in the tool results.

# Logging

The agent uses structured logging (slog) to provide detailed execution information:
  - Configuration details on initialization
  - Task execution progress and tool usage
  - Error conditions with context
  - Tool execution results and messages

Logs can be configured through the WithLogger option to capture different levels
of detail as needed.

# Thread Safety

The agent is designed to be thread-safe and can handle multiple concurrent tasks.
Each task execution gets its own context and can be cancelled independently.
*/
package agent
