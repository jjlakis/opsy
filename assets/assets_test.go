package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRenderAgentSystemPrompt(t *testing.T) {
	t.Run("renders with valid data", func(t *testing.T) {
		data := &AgentSystemPromptData{
			Shell: "/bin/bash",
		}
		result, err := RenderAgentSystemPrompt(data)
		require.NoError(t, err)
		assert.Contains(t, result, "/bin/bash")
		assert.NotEmpty(t, result)
	})

	t.Run("handles empty shell", func(t *testing.T) {
		data := &AgentSystemPromptData{}
		result, err := RenderAgentSystemPrompt(data)
		require.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("handles nil data", func(t *testing.T) {
		_, err := RenderAgentSystemPrompt(nil)
		assert.Error(t, err)
	})
}

func TestRenderToolSystemPrompt(t *testing.T) {
	t.Run("renders with valid data", func(t *testing.T) {
		data := &ToolSystemPromptData{
			Shell:      "/bin/bash",
			Name:       "test-tool",
			Executable: "/usr/bin/test",
			Rules:      []string{"rule1", "rule2"},
		}
		result, err := RenderToolSystemPrompt(data)
		require.NoError(t, err)
		assert.Contains(t, result, "/bin/bash")
		assert.Contains(t, result, "test-tool")
		assert.Contains(t, result, "/usr/bin/test")
		assert.Contains(t, result, "rule1")
		assert.Contains(t, result, "rule2")
		assert.NotEmpty(t, result)
	})

	t.Run("handles empty fields", func(t *testing.T) {
		data := &ToolSystemPromptData{}
		result, err := RenderToolSystemPrompt(data)
		require.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("handles nil data", func(t *testing.T) {
		_, err := RenderToolSystemPrompt(nil)
		assert.Error(t, err)
	})

	t.Run("handles empty rules", func(t *testing.T) {
		data := &ToolSystemPromptData{
			Shell:      "/bin/bash",
			Name:       "test-tool",
			Executable: "/usr/bin/test",
		}
		result, err := RenderToolSystemPrompt(data)
		require.NoError(t, err)
		assert.NotEmpty(t, result)
	})
}

func TestRenderToolUserPrompt(t *testing.T) {
	t.Run("renders with valid data", func(t *testing.T) {
		data := &ToolUserPromptData{
			Task: "test task",
			Params: map[string]any{
				"param1": "value1",
				"param2": 42,
			},
			Context: map[string]string{
				"ctx1": "value1",
				"ctx2": "value2",
			},
			WorkingDirectory: "/test/dir",
		}
		result, err := RenderToolUserPrompt(data)
		require.NoError(t, err)
		assert.Contains(t, result, "test task")
		assert.Contains(t, result, "param1")
		assert.Contains(t, result, "value1")
		assert.Contains(t, result, "ctx1")
		assert.Contains(t, result, "/test/dir")
		assert.NotEmpty(t, result)
	})

	t.Run("handles empty fields", func(t *testing.T) {
		data := &ToolUserPromptData{}
		result, err := RenderToolUserPrompt(data)
		require.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("handles nil data", func(t *testing.T) {
		_, err := RenderToolUserPrompt(nil)
		assert.Error(t, err)
	})

	t.Run("handles empty maps", func(t *testing.T) {
		data := &ToolUserPromptData{
			Task:             "test task",
			WorkingDirectory: "/test/dir",
		}
		result, err := RenderToolUserPrompt(data)
		require.NoError(t, err)
		assert.NotEmpty(t, result)
	})
}

func TestEmbeddedFS(t *testing.T) {
	t.Run("themes fs is accessible", func(t *testing.T) {
		entries, err := Themes.ReadDir(ThemeDir)
		require.NoError(t, err)
		assert.NotEmpty(t, entries)
	})

	t.Run("tools fs is accessible", func(t *testing.T) {
		entries, err := Tools.ReadDir(ToolsDir)
		require.NoError(t, err)
		assert.NotEmpty(t, entries)
	})
}
