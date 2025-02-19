package toolmanager

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os/exec"
	"time"

	"github.com/datolabs-io/sredo/internal/config"
	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
	"golang.org/x/exp/maps"
)

// Tool is the interface for a tool.
type Tool interface {
	// GetName returns the name of the tool.
	GetName() string
	// GetDisplayName returns the display name of the tool.
	GetDisplayName() string
	// GetDescription returns the description of the tool.
	GetDescription() string
	// GetInputSchema returns the input schema of the tool.
	GetInputSchema() *jsonschema.Schema
	// Execute executes the tool.
	Execute(inputs map[string]any, ctx context.Context) (*ToolOutput, error)
}

// Tool is a tool that can be used by the agent.
type tool struct {
	// name is the name of the tool.
	name string
	// config is the config for the tool.
	config *config.ToolsConfiguration
	// definition is the definition of the tool.
	definition toolDefinition
	// logger is the logger for the tool.
	logger *slog.Logger
	// inputSchema is the input schema of the tool.
	inputSchema *jsonschema.Schema
}

// toolDefinition is the definition of a tool.
type toolDefinition struct {
	// DisplayName is the name of the tool as it will be displayed in the UI.
	DisplayName string `yaml:"display_name"`
	// Description is the description of the tool as it will be displayed in the UI.
	Description string `yaml:"description"`
	// SystemPrompt is the system prompt to use when the tool is selected.
	SystemPrompt string `yaml:"system_prompt"`
	// Inputs is the inputs for the tool.
	Inputs map[string]toolInput `yaml:"inputs"`
	// Executable is the executable to use to execute the tool.
	Executable string `yaml:"executable,omitempty"`
}

// toolInput is the definition of an input for a tool.
type toolInput struct {
	// Type is the type of the input.
	Type string `yaml:"type"`
	// Description is the description of the input.
	Description string `yaml:"description"`
	// Default is the default value for the input.
	Default string `yaml:"default"`
	// Examples are examples of the input.
	Examples []any `yaml:"examples"`
	// Optional is whether the input is optional.
	Optional bool `yaml:"optional"`
}

// ToolOutput is the output of a tool.
type ToolOutput struct {
	// Tool is the name of the tool that executed the task.
	Tool string `json:"tool"`
	// Result is the result from the tool execution.
	Result any `json:"result,omitempty"`
	// IsError indicates if the tool execution resulted in an error.
	IsError bool `json:"is_error"`
	// ExecutedCommand is the command that was executed.
	ExecutedCommand *Command `json:"executed_command,omitempty"`
}

const (
	// ErrInvalidInputType is the error returned when an input parameter has an invalid type.
	ErrInvalidToolInputType = "invalid input type"
	// ErrToolMissingDisplayName is the error returned when a tool is missing a display name.
	ErrToolMissingDisplayName = "missing tool display name"
	// ErrToolMissingDescription is the error returned when a tool is missing a description.
	ErrToolMissingDescription = "missing tool description"
	// ErrToolInputMissingType is the error returned when a tool input is missing a type.
	ErrToolInputMissingType = "missing tool input type"
	// ErrToolInputMissingDescription is the error returned when a tool input is missing a description.
	ErrToolInputMissingDescription = "missing tool input description"
	// ErrToolExecutableNotFound is the error returned when a tool executable is not found.
	ErrToolExecutableNotFound = "tool executable not found"

	// InputTask is the input parameter for the task to complete.
	InputTask = "task"
	// InputWorkingDirectory is the input parameter for the working directory to use for the tool.
	InputWorkingDirectory = "working_directory"
	// InputCommand is the input parameter for the command to execute.
	InputCommand = "command"
)

// newTool creates a new tool.
func newTool(n string, def toolDefinition, prompt string, logger *slog.Logger, cfg *config.ToolsConfiguration) *tool {
	logger = logger.WithGroup("tool").With("name", n).With("display_name", def.DisplayName).
		With("description", def.Description).With("executable", def.Executable)

	tool := &tool{
		definition:  def,
		inputSchema: generateInputSchema(appendCommonInputs(def.Inputs)),
		config:      cfg,
		logger:      logger,
		name:        n,
	}

	tool.definition.SystemPrompt = fmt.Sprintf("%s\n\n%s", def.SystemPrompt, prompt)
	tool.logger.Debug("Tool loaded.")

	return tool
}

// GetName returns the name of the tool.
func (t *tool) GetName() string {
	return t.name
}

// GetDisplayName returns the display name of the tool.
func (t *tool) GetDisplayName() string {
	return t.definition.DisplayName
}

// GetDescription returns the description of the tool.
func (t *tool) GetDescription() string {
	return t.definition.Description
}

// GetInputSchema returns the input schema of the tool.
func (t *tool) GetInputSchema() *jsonschema.Schema {
	return t.inputSchema
}

// Execute executes the tool.
func (t *tool) Execute(inputs map[string]any, ctx context.Context) (*ToolOutput, error) {
	t.logger.With("inputs", inputs).Info("Executing tool.")

	ctx, cancel := context.WithTimeout(ctx, time.Duration(t.config.Timeout))
	defer cancel()

	// TODO(t-dabasinskas): Implement the tool's logic here.
	return nil, nil
}

// appendCommonInputs appends the common tool inputs to the tool's inputs.
func appendCommonInputs(inputs map[string]toolInput) map[string]toolInput {
	allInputs := map[string]toolInput{
		InputTask: {
			Type:        "string",
			Description: "Task (without parameters) the user wants to complete with the tool.",
			Examples: []any{
				"Clone the repository",
				"Get deployments",
			},
			Default:  "",
			Optional: false,
		},
		InputWorkingDirectory: {
			Type:        "string",
			Description: "Working directory to use for the tool.",
			Examples: []any{
				"~/projects/my-project",
				"/tmp",
			},
			Default:  ".",
			Optional: true,
		},
	}

	maps.Copy(allInputs, inputs)
	return allInputs
}

// GenerateInputSchema generates a JSON schema for the tool's inputs.
func generateInputSchema(inputs map[string]toolInput) *jsonschema.Schema {
	required := make([]string, 0)
	properties := orderedmap.New[string, *jsonschema.Schema]()

	for name, input := range inputs {
		properties.Set(name, &jsonschema.Schema{
			Type:        input.Type,
			Description: input.Description,
			Default:     input.Default,
			Examples:    input.Examples,
		})

		if !input.Optional {
			required = append(required, name)
		}
	}

	schema := &jsonschema.Schema{
		Properties: properties,
		Required:   required,
		Type:       "object",
	}

	return schema
}

// validateToolDefinition validates a tool definition.
func validateToolDefinition(def *toolDefinition) error {
	if def.DisplayName == "" {
		return errors.New(ErrToolMissingDisplayName)
	}
	if def.Description == "" {
		return errors.New(ErrToolMissingDescription)
	}

	for name, input := range def.Inputs {
		if input.Type == "" {
			return fmt.Errorf("%s: %q", ErrToolInputMissingType, name)
		}
		if input.Description == "" {
			return fmt.Errorf("%s: %q", ErrToolInputMissingDescription, name)
		}
	}

	if def.Executable != "" {
		if _, err := exec.LookPath(def.Executable); err != nil {
			return fmt.Errorf("%s: %q", ErrToolExecutableNotFound, def.Executable)
		}
	}

	return nil
}
