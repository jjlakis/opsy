package toolmanager

import (
	"context"
	"io"
	"log/slog"
	"testing"

	"github.com/datolabs-io/sredo/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewTool tests the creation of a new tool.
func TestNewTool(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(nil, nil))
	cfg := &config.ToolsConfiguration{
		Timeout: 120,
		Exec: config.ExecToolConfiguration{
			Timeout: 60,
			Shell:   "/bin/bash",
		},
	}
	definition := toolDefinition{
		DisplayName:  "Test Tool",
		Description:  "Test Description",
		SystemPrompt: "Test Prompt",
		Inputs:       make(map[string]toolInput),
	}

	tool := newTool("test", definition, "Default Prompt", logger, cfg)

	t.Run("initializes with correct values", func(t *testing.T) {
		assert.Equal(t, "test", tool.name)
		assert.Equal(t, "Test Tool", tool.GetDisplayName())
		assert.Equal(t, "Test Description", tool.GetDescription())
		assert.Equal(t, "Test Prompt\n\nDefault Prompt", tool.definition.SystemPrompt)
		assert.NotNil(t, tool.GetInputSchema())
		assert.Equal(t, cfg, tool.config)
	})
}

// TestToolGetters tests the getter methods of Tool.
func TestToolGetters(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(nil, nil))
	cfg := &config.ToolsConfiguration{
		Timeout: 120,
		Exec: config.ExecToolConfiguration{
			Timeout: 60,
			Shell:   "/bin/bash",
		},
	}
	inputs := map[string]toolInput{
		"test_input": {
			Type:        "string",
			Description: "Test Input",
			Default:     "default",
			Examples:    []any{"example"},
			Optional:    false,
		},
	}

	definition := toolDefinition{
		DisplayName:  "Display Name",
		Description:  "Description",
		SystemPrompt: "System Prompt",
		Inputs:       inputs,
	}

	tool := newTool("test", definition, "", logger, cfg)

	t.Run("GetDisplayName returns correct value", func(t *testing.T) {
		assert.Equal(t, "Display Name", tool.GetDisplayName())
	})

	t.Run("GetDescription returns correct value", func(t *testing.T) {
		assert.Equal(t, "Description", tool.GetDescription())
	})

	t.Run("GetInputSchema returns valid schema", func(t *testing.T) {
		schema := tool.GetInputSchema()
		require.NotNil(t, schema)

		// Verify schema properties
		prop, ok := schema.Properties.Get("test_input")
		require.True(t, ok)
		assert.Equal(t, "string", prop.Type)
		assert.Equal(t, "Test Input", prop.Description)
		assert.Equal(t, "default", prop.Default)
		assert.Equal(t, []any{"example"}, prop.Examples)
	})
}

// TestGenerateInputSchema tests the schema generation for tool inputs.
func TestGenerateInputSchema(t *testing.T) {
	inputs := map[string]toolInput{
		"required_input": {
			Type:        "string",
			Description: "Required Input",
			Default:     "default",
			Examples:    []any{"example1", "example2"},
			Optional:    false,
		},
		"optional_input": {
			Type:        "string",
			Description: "Optional Input",
			Default:     "",
			Examples:    []any{"optional"},
			Optional:    true,
		},
	}

	schema := generateInputSchema(inputs)

	t.Run("generates correct schema structure", func(t *testing.T) {
		assert.Equal(t, "object", schema.Type)
		assert.Contains(t, schema.Required, "required_input")
		assert.NotContains(t, schema.Required, "optional_input")
	})

	t.Run("includes all inputs in properties", func(t *testing.T) {
		requiredProp, ok := schema.Properties.Get("required_input")
		require.True(t, ok)
		assert.Equal(t, "Required Input", requiredProp.Description)
		assert.Equal(t, "default", requiredProp.Default)
		assert.Equal(t, []any{"example1", "example2"}, requiredProp.Examples)

		optionalProp, ok := schema.Properties.Get("optional_input")
		require.True(t, ok)
		assert.Equal(t, "Optional Input", optionalProp.Description)
		assert.Equal(t, "", optionalProp.Default)
		assert.Equal(t, []any{"optional"}, optionalProp.Examples)
	})
}

// TestToolExecute tests the Execute method of Tool.
func TestToolExecute(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	cfg := &config.ToolsConfiguration{
		Timeout: 120,
		Exec: config.ExecToolConfiguration{
			Timeout: 60,
			Shell:   "/bin/bash",
		},
	}
	tool := newTool("test", toolDefinition{
		DisplayName: "Test Tool",
		Description: "Test Description",
		Inputs:      map[string]toolInput{}, // Initialize with empty inputs
	}, "", logger, cfg)

	t.Run("returns nil result", func(t *testing.T) {
		input := map[string]any{"test": "value"}
		output, err := tool.Execute(input, context.Background())
		require.NoError(t, err)
		assert.Nil(t, output)
	})
}

// TestAppendCommonInputs tests the appendCommonInputs function.
func TestAppendCommonInputs(t *testing.T) {
	customInputs := map[string]toolInput{
		"custom_input": {
			Type:        "string",
			Description: "Custom input",
			Default:     "default",
			Examples:    []any{"example"},
			Optional:    false,
		},
	}

	allInputs := appendCommonInputs(customInputs)

	t.Run("includes common inputs", func(t *testing.T) {
		// Check task input
		taskInput, ok := allInputs[InputTask]
		require.True(t, ok, "task input should be present")
		assert.Equal(t, "string", taskInput.Type)
		assert.Equal(t, "Task (without parameters) the user wants to complete with the tool.", taskInput.Description)
		assert.False(t, taskInput.Optional)

		// Check working directory input
		wdInput, ok := allInputs[InputWorkingDirectory]
		require.True(t, ok, "working directory input should be present")
		assert.Equal(t, "string", wdInput.Type)
		assert.Equal(t, "Working directory to use for the tool.", wdInput.Description)
		assert.Equal(t, ".", wdInput.Default)
		assert.True(t, wdInput.Optional)
	})

	t.Run("preserves custom inputs", func(t *testing.T) {
		customInput, ok := allInputs["custom_input"]
		require.True(t, ok, "custom input should be present")
		assert.Equal(t, "string", customInput.Type)
		assert.Equal(t, "Custom input", customInput.Description)
		assert.Equal(t, "default", customInput.Default)
		assert.Equal(t, []any{"example"}, customInput.Examples)
		assert.False(t, customInput.Optional)
	})
}

// TestValidateToolDefinition tests the validateToolDefinition function.
func TestValidateToolDefinition(t *testing.T) {
	t.Run("validates valid tool definition", func(t *testing.T) {
		def := &toolDefinition{
			DisplayName:  "Valid Tool",
			Description:  "Valid Description",
			SystemPrompt: "Valid Prompt",
			Inputs: map[string]toolInput{
				"input1": {
					Type:        "string",
					Description: "Valid Input",
					Default:     "default",
					Examples:    []any{"example"},
					Optional:    false,
				},
			},
		}
		err := validateToolDefinition(def)
		assert.NoError(t, err)
	})

	t.Run("validates empty tool definition", func(t *testing.T) {
		def := &toolDefinition{}
		err := validateToolDefinition(def)
		assert.ErrorContains(t, err, "display_name is required")

		def.DisplayName = "Tool"
		err = validateToolDefinition(def)
		assert.ErrorContains(t, err, "description is required")
	})

	t.Run("validates input fields", func(t *testing.T) {
		def := &toolDefinition{
			DisplayName: "Tool",
			Description: "Description",
			Inputs: map[string]toolInput{
				"input1": {},
			},
		}
		err := validateToolDefinition(def)
		assert.ErrorContains(t, err, `type is required for input "input1"`)

		def.Inputs["input1"] = toolInput{Type: "string"}
		err = validateToolDefinition(def)
		assert.ErrorContains(t, err, `description is required for input "input1"`)

		def.Inputs["input1"] = toolInput{
			Type:        "string",
			Description: "Description",
		}
		err = validateToolDefinition(def)
		assert.NoError(t, err)
	})

	t.Run("allows empty inputs", func(t *testing.T) {
		def := &toolDefinition{
			DisplayName: "Tool",
			Description: "Description",
		}
		err := validateToolDefinition(def)
		assert.NoError(t, err)
	})
}

// TestToolInterfaceCompliance tests that tool implementations comply with the Tool interface.
func TestToolInterfaceCompliance(t *testing.T) {
	// Test regular tool
	var _ Tool = &tool{}

	// Test exec tool
	var _ Tool = (*execTool)(nil)

	// Create and test a concrete tool instance
	logger := slog.New(slog.NewTextHandler(nil, nil))
	cfg := &config.ToolsConfiguration{}
	def := toolDefinition{
		DisplayName: "Test Tool",
		Description: "Test Description",
		Inputs: map[string]toolInput{
			"test": {
				Type:        "string",
				Description: "Test Input",
			},
		},
	}

	tool := newTool("test", def, "", logger, cfg)

	t.Run("implements all interface methods", func(t *testing.T) {
		assert.NotPanics(t, func() {
			_ = tool.GetName()
			_ = tool.GetDisplayName()
			_ = tool.GetDescription()
			_ = tool.GetInputSchema()
			_, _ = tool.Execute(nil, context.Background())
		})
	})

	t.Run("exec tool implements all interface methods", func(t *testing.T) {
		execTool := newExecTool(logger, cfg)
		assert.NotPanics(t, func() {
			_ = execTool.GetName()
			_ = execTool.GetDisplayName()
			_ = execTool.GetDescription()
			_ = execTool.GetInputSchema()
			_, _ = execTool.Execute(nil, context.Background())
		})
	})
}
