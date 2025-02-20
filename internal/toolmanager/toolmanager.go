package toolmanager

import (
	"context"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/datolabs-io/sredo/assets"
	"github.com/datolabs-io/sredo/internal/config"
	"gopkg.in/yaml.v3"
)

const (
	// ErrLoadingTools is the error message for failed to load tools.
	ErrLoadingTools = "failed to load tools"
	// ErrLoadingTool is the error message for failed to load a specific tool.
	ErrLoadingTool = "failed to load tool"
	// ErrParsingTool is the error message for failed to parse a tool.
	ErrParsingTool = "failed to parse tool"
	// ErrToolNotFound is the error message for a tool not found.
	ErrToolNotFound = "tool not found"
	// ErrInvalidToolDefinition is the error message for an invalid tool definition.
	ErrInvalidToolDefinition = "invalid tool definition"

	// toolsDir is the directory containing the tools.
	toolsDir = "tools"

	// commonToolSystemPrompt is the default system prompt added for all tools.
	commonToolSystemPrompt = ``
)

// Manager is the interface for the tool manager.
type Manager interface {
	// LoadTools loads the tools from the tool manager.
	LoadTools() error
	// GetTools returns all tools.
	GetTools() map[string]Tool
	// GetTool returns a tool by name.
	GetTool(name string) (Tool, error)
}

// ToolManager is the tool manager.
type ToolManager struct {
	cfg    config.Configuration
	logger *slog.Logger
	ctx    context.Context
	fs     fs.FS
	dir    string
	tools  map[string]Tool
}

// Option is a function that modifies the tool manager.
type Option func(*ToolManager)

// New creates a new tool manager.
func New(opts ...Option) *ToolManager {
	tm := &ToolManager{
		cfg:    config.New().GetConfig(),
		logger: slog.New(slog.DiscardHandler),
		ctx:    context.Background(),
		fs:     assets.Tools,
		dir:    toolsDir,
		tools:  make(map[string]Tool),
	}

	for _, opt := range opts {
		opt(tm)
	}

	tm.logger.WithGroup("config").With("directory", tm.dir).Debug("Tool manager initialized.")

	return tm
}

// WithConfig sets the configuration for the tool manager.
func WithConfig(cfg config.Configuration) Option {
	return func(tm *ToolManager) {
		tm.cfg = cfg
	}
}

// WithLogger sets the logger for the tool manager.
func WithLogger(logger *slog.Logger) Option {
	return func(tm *ToolManager) {
		tm.logger = logger.With("component", "toolmanager")
	}
}

// WithDirectory sets the directory for the tool manager.
func WithDirectory(dir string) Option {
	return func(tm *ToolManager) {
		tm.fs = os.DirFS(dir)
		tm.dir = "."
	}
}

// WithContext sets the context for the tool manager.
func WithContext(ctx context.Context) Option {
	return func(tm *ToolManager) {
		tm.ctx = ctx
	}
}

// LoadTools loads the tools from the tool manager.
func (tm *ToolManager) LoadTools() error {
	toolFiles, err := fs.ReadDir(tm.fs, tm.dir)
	if err != nil {
		return fmt.Errorf("%s: %v", ErrLoadingTools, err)
	}

	// Exec tool is a special tool which we always statically load.
	tm.tools[ExecToolName] = newExecTool(tm.logger, &tm.cfg.Tools)

	for _, toolFile := range toolFiles {
		if toolFile.IsDir() {
			continue
		}

		name := strings.TrimSuffix(toolFile.Name(), filepath.Ext(toolFile.Name()))
		tool, err := tm.loadTool(name, toolFile)
		if err != nil {
			tm.logger.With("tool.name", name).With("filename", toolFile.Name()).With("error", err).
				Error("Failed to load the tool.")
			continue
		}

		tm.tools[name] = tool
	}

	tm.logger.With("tools.count", len(tm.tools)).Debug("Tools loaded.")

	return nil
}

// loadTool loads a tool from a file.
func (tm *ToolManager) loadTool(name string, toolFile fs.DirEntry) (*tool, error) {
	contents, err := fs.ReadFile(tm.fs, filepath.Join(tm.dir, toolFile.Name()))
	if err != nil {
		return nil, fmt.Errorf("%s: %v", ErrLoadingTool, err)
	}

	var definition toolDefinition
	if err := yaml.Unmarshal(contents, &definition); err != nil {
		return nil, fmt.Errorf("%s: %v", ErrParsingTool, err)
	}

	if err := validateToolDefinition(&definition); err != nil {
		return nil, fmt.Errorf("%s: %s: %v", ErrInvalidToolDefinition, name, err)
	}

	return newTool(name, definition, commonToolSystemPrompt, tm.logger, &tm.cfg.Tools), nil
}

// GetTools returns all tools.
func (tm *ToolManager) GetTools() map[string]Tool {
	return tm.tools
}

// GetTool returns a tool by name.
func (tm *ToolManager) GetTool(name string) (Tool, error) {
	tool, ok := tm.tools[name]
	if !ok {
		return nil, fmt.Errorf("%s: %v", ErrToolNotFound, name)
	}

	return tool, nil
}
