package toolmanager

import (
	"context"
	"encoding/json"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewTool tests the creation of a new tool.
func TestNewTool(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(nil, nil))
	definition := toolDefinition{
		DisplayName:  "Test Tool",
		Description:  "Test Description",
		SystemPrompt: "Test Prompt",
		Inputs:       &toolInputs{},
	}

	tool := newTool("test", definition, "Default Prompt", logger)

	t.Run("initializes with correct values", func(t *testing.T) {
		assert.Equal(t, "test", tool.name)
		assert.Equal(t, "Test Tool", tool.GetDisplayName())
		assert.Equal(t, "Test Description", tool.GetDescription())
		assert.Equal(t, "Test Prompt\n\nDefault Prompt", tool.definition.SystemPrompt)
		assert.NotNil(t, tool.GetInputSchema())
	})
}

// TestToolGetters tests the getter methods of Tool.
func TestToolGetters(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(nil, nil))
	inputs := &toolInputs{
		"test_input": {
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

	tool := newTool("test", definition, "", logger)

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
	inputs := &toolInputs{
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
	logger := slog.New(slog.NewTextHandler(nil, nil))
	tool := newTool("test", toolDefinition{
		DisplayName: "Test Tool",
		Description: "Test Description",
		Inputs:      &toolInputs{}, // Initialize with empty inputs
	}, "", logger)

	t.Run("returns empty result", func(t *testing.T) {
		input := json.RawMessage(`{"test": "value"}`)
		result, err := tool.Execute(input, context.Background())
		require.NoError(t, err)
		assert.Nil(t, output)
	})
}
