package toolmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
	"golang.org/x/exp/maps"
)

// Tooling is the interface for a tool.
type Tooling interface {
	// GetName returns the name of the tool.
	GetName() string
	// GetDisplayName returns the display name of the tool.
	GetDisplayName() string
	// GetDescription returns the description of the tool.
	GetDescription() string
	// GetInputSchema returns the input schema of the tool.
	GetInputSchema() *jsonschema.Schema
	// Execute executes the tool.
	Execute(input json.RawMessage, ctx context.Context) (string, error)
}

// Tool is a tool that can be used by the agent.
type Tool struct {
	// name is the name of the tool.
	name string
	// definition is the definition of the tool.
	definition toolDefinition
	// logger is the logger for the tool.
	logger *slog.Logger
	// inputSchema is the input schema of the tool.
	inputSchema *jsonschema.Schema
}

// toolInputs is an ordered map of tool inputs.
type toolInputs map[string]toolInput

// toolDefinition is the definition of a tool.
type toolDefinition struct {
	// DisplayName is the name of the tool as it will be displayed in the UI.
	DisplayName string `yaml:"display_name"`
	// Description is the description of the tool as it will be displayed in the UI.
	Description string `yaml:"description"`
	// SystemPrompt is the system prompt to use when the tool is selected.
	SystemPrompt string `yaml:"system_prompt"`
	// Inputs is the inputs for the tool.
	Inputs *toolInputs
}

// toolInput is the definition of an input for a tool.
type toolInput struct {
	// Description is the description of the input.
	Description string `yaml:"description"`
	// Default is the default value for the input.
	Default string `yaml:"default"`
	// Examples are examples of the input.
	Examples []any `yaml:"examples"`
	// Optional is whether the input is optional.
	Optional bool `yaml:"optional"`
}

// commonToolInputs is the set of inputs that are common to all tools.
var commonToolInputs = toolInputs{
	"task": {
		Description: "Task (without parameters) the user wants to complete with the tool.",
		Examples: []any{
			"Clone the repository",
			"Get deployments",
		},
		Default:  "",
		Optional: false,
	},
	"working_directory": {
		Description: "Working directory to use for the tool.",
		Examples: []any{
			"~/projects/my-project",
			"/tmp",
		},
		Default:  ".",
		Optional: true,
	},
}

// newTool creates a new tool.
func newTool(name string, definition toolDefinition, defaultPrompt string, logger *slog.Logger) Tool {
	tool := &Tool{
		definition:  definition,
		inputSchema: generateInputSchema(definition),
		logger:      logger,
		name:        name,
	}

	tool.definition.SystemPrompt = fmt.Sprintf("%s\n\n%s", definition.SystemPrompt, defaultPrompt)

	tool.logger.With("name", name).With("display_name", definition.DisplayName).Debug("Tool loaded.")

	return *tool
}

// GetDisplayName returns the display name of the tool.
func (t *Tool) GetDisplayName() string {
	return t.definition.DisplayName
}

// GetDescription returns the description of the tool.
func (t *Tool) GetDescription() string {
	return t.definition.Description
}

// GetInputSchema returns the input schema of the tool.
func (t *Tool) GetInputSchema() *jsonschema.Schema {
	return t.inputSchema
}

// Execute executes the tool.
func (t *Tool) Execute(input json.RawMessage, ctx context.Context) (string, error) {
	// TODO: Implement the tool's logic here.
	return nil, nil
}

// appendCommonInputs appends the common tool inputs to the tool's inputs.
func appendCommonInputs(inputs map[string]toolInput) map[string]toolInput {
	allInputs := map[string]toolInput{
		InputTask: {
			Description: "Task (without parameters) the user wants to complete with the tool.",
			Examples: []any{
				"Clone the repository",
				"Get deployments",
			},
			Default:  "",
			Optional: false,
		},
		InputWorkingDirectory: {
			Description: "Working directory to use for the tool.",
			Examples: []any{
				"~/projects/my-project",
				"/tmp",
			},
			Default:  ".",
			Optional: true,
		},
	}

	for name, input := range inputs {
		allInputs[name] = input
	}

	return inputs
}

// GenerateInputSchema generates a JSON schema for the tool's inputs.
func generateInputSchema(inputs map[string]toolInput) *jsonschema.Schema {
	required := make([]string, 0)
	properties := orderedmap.New[string, *jsonschema.Schema]()

	for name, input := range *definition.Inputs {
		properties.Set(name, &jsonschema.Schema{
			Type:        "string",
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
