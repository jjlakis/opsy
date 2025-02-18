package toolmanager

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/datolabs-io/sredo/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNew tests the creation of a new tool manager with various options.
func TestNew(t *testing.T) {
	t.Run("creates default tool manager", func(t *testing.T) {
		tm := New()
		assert.NotNil(t, tm)
		assert.Equal(t, "tools", tm.dir)
		assert.NotNil(t, tm.tools)
	})

	t.Run("creates tool manager with options", func(t *testing.T) {
		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		cfg := config.New().GetConfig()
		ctx := context.Background()

		tm := New(
			WithLogger(logger),
			WithConfig(cfg),
			WithContext(ctx),
			WithDirectory("testdata"),
		)

		assert.NotNil(t, tm)
		assert.Equal(t, ".", tm.dir)
		assert.Equal(t, cfg, tm.cfg)
		assert.Equal(t, ctx, tm.ctx)
	})
}

// TestLoadTools tests loading tools from the filesystem.
func TestLoadTools(t *testing.T) {
	t.Run("loads valid tools", func(t *testing.T) {
		tm := New(WithDirectory("testdata"))
		err := tm.LoadTools()
		require.NoError(t, err)

		tools := tm.GetTools()
		assert.Len(t, tools, 1) // Should only load test_tool.yaml, not invalid_tool.yaml

		tool, err := tm.GetTool("test_tool")
		require.NoError(t, err)
		assert.Equal(t, "Test Tool", tool.GetDisplayName())
		assert.Equal(t, "A tool for testing purposes", tool.GetDescription())
	})

	t.Run("handles invalid tool definitions", func(t *testing.T) {
		// The invalid tool should be skipped during loading
		tm := New(WithDirectory("testdata"))
		err := tm.LoadTools()
		require.NoError(t, err)

		_, err = tm.GetTool("invalid_tool")
		assert.ErrorContains(t, err, ErrToolNotFound)
	})

	t.Run("handles non-existent directory", func(t *testing.T) {
		tm := New(WithDirectory("nonexistent"))
		err := tm.LoadTools()
		assert.ErrorContains(t, err, ErrLoadingTools)
	})
}

// TestGetTool tests retrieving specific tools.
func TestGetTool(t *testing.T) {
	tm := New(WithDirectory("testdata"))
	require.NoError(t, tm.LoadTools())

	t.Run("gets existing tool", func(t *testing.T) {
		tool, err := tm.GetTool("test_tool")
		require.NoError(t, err)
		assert.Equal(t, "Test Tool", tool.GetDisplayName())
	})

	t.Run("returns error for non-existent tool", func(t *testing.T) {
		_, err := tm.GetTool("nonexistent")
		assert.ErrorContains(t, err, ErrToolNotFound)
	})
}

// TestGetTools tests retrieving all tools.
func TestGetTools(t *testing.T) {
	tm := New(WithDirectory("testdata"))
	require.NoError(t, tm.LoadTools())

	tools := tm.GetTools()
	assert.Len(t, tools, 1)
	assert.Equal(t, "Test Tool", tools[0].GetDisplayName())
}
