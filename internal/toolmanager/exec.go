package toolmanager

import (
	"context"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/datolabs-io/sredo/internal/config"
	"github.com/invopop/jsonschema"
)

// ExecTool is the tool for executing commands.
type execTool tool

// ExecToolName is the name of the exec tool.
const ExecToolName = "exec"

// Command is the command that was executed.
type Command struct {
	// Command is the command that was executed.
	Command string
	// WorkingDirectory is the working directory of the command.
	WorkingDirectory string
	// ExitCode is the exit code of the command.
	ExitCode int
	// Output is the output of the command.
	Output string
	// StartedAt is the time the command started.
	StartedAt time.Time
	// CompletedAt is the time the command completed.
	CompletedAt time.Time
}

// newExecTool creates a new exec tool.
func newExecTool(logger *slog.Logger, cfg *config.ToolsConfiguration) *execTool {
	definition := toolDefinition{
		DisplayName: "Exec",
		Description: "Executes the provided shell command.",
		Inputs: map[string]toolInput{
			InputCommand: {
				Description: "The shell command, including all the arguments, to execute",
				Type:        "string",
				Examples: []any{
					"ls -l | grep 'myfile'",
					"git status",
					"curl -X GET https://api.example.com/data",
				},
			},
		},
	}

	return (*execTool)(newTool(ExecToolName, definition, commonToolSystemPrompt, logger, cfg))
}

// GetName returns the name of the tool.
func (t *execTool) GetName() string {
	return (*tool)(t).name
}

// GetDisplayName returns the display name of the tool.
func (t *execTool) GetDisplayName() string {
	return (*tool)(t).GetDisplayName()
}

// GetDescription returns the description of the tool.
func (t *execTool) GetDescription() string {
	return (*tool)(t).GetDescription()
}

// GetInputSchema returns the input schema of the tool.
func (t *execTool) GetInputSchema() *jsonschema.Schema {
	return (*tool)(t).GetInputSchema()
}

// Execute executes the tool.
func (t *execTool) Execute(inputs map[string]any, ctx context.Context) (*ToolOutput, error) {
	command, ok := inputs[InputCommand].(string)
	if !ok {
		return nil, fmt.Errorf("%s: %s", ErrInvalidToolInputType, InputCommand)
	}

	workingDirectory, ok := inputs[InputWorkingDirectory].(string)
	if !ok {
		return nil, fmt.Errorf("%s: %s", ErrInvalidToolInputType, InputWorkingDirectory)
	}

	ctx, cancel := context.WithTimeout(ctx, t.getTimeout())
	defer cancel()

	cmd := exec.CommandContext(ctx, t.config.Exec.Shell, "-c", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Dir = workingDirectory
	cmd.Stdin = nil
	startedAt := time.Now()

	logger := t.logger.With("command", cmd.String()).With("working_directory", workingDirectory)
	logger.Debug("Executing command.")

	toolOutput, err := cmd.CombinedOutput()
	output := &ToolOutput{
		Tool:    t.GetName(),
		IsError: false,
		Result:  string(toolOutput),
		ExecutedCommand: &Command{
			Command:          command,
			WorkingDirectory: workingDirectory,
			ExitCode:         cmd.ProcessState.ExitCode(),
			StartedAt:        startedAt,
			CompletedAt:      time.Now(),
		},
	}

	if toolOutput != nil {
		output.ExecutedCommand.Output = strings.TrimSpace(string(toolOutput))
		output.Result = output.ExecutedCommand.Output
	}

	if err != nil {
		logger.With("error", err).With("exit_code", cmd.ProcessState.ExitCode()).Error("Command execution failed.")
		output.IsError = true
	}

	return output, err
}

// getTimeout returns the timeout for the Exec tool.
func (t *execTool) getTimeout() time.Duration {
	timeout := t.config.Timeout
	if t.config.Exec.Timeout > 0 {
		timeout = t.config.Exec.Timeout
	}

	return time.Duration(timeout) * time.Second
}
