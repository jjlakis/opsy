package agent

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/datolabs-io/sredo/internal/config"
	"github.com/datolabs-io/sredo/internal/toolmanager"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

//go:embed prompt.mdx
var agentPrompt string

const (
	// ErrNoRunOptions is the error returned when no run options are provided.
	ErrNoRunOptions = "no run options provided"
	// ErrNoTaskProvided is the error returned when no task is provided.
	ErrNoTaskProvided = "no task provided"

	// StatusReady is the status of the agent when it is ready to run.
	StatusReady = "Ready"
	// StatusRunning is the status of the agent when it is running.
	StatusRunning = "Running"
	// StatusFinished is the status of the agent when it has finished.
	StatusFinished = "Finished"
	// StatusError is the status of the agent when it has encountered an error.
	StatusError = "Error"
)

// Status is the status of the agent.
type Status string

// AgentRunner is an interface that defines the methods for an agent.
type AgentRunner interface {
	Run(task string, prompt string, tools []toolmanager.Tool, ctx context.Context) (string, error)
}

// Agent is a struct that contains the state of the agent.
type Agent struct {
	client        *anthropic.Client
	ctx           context.Context
	cfg           config.Configuration
	logger        *slog.Logger
	communication *Communication
}

// Message is a struct that contains a message from the agent.
type Message struct {
	// Tool is the name of the tool that sent the message.
	Tool string
	// Message is the message from the tool.
	Message string
	// Timestamp is the timestamp when the message was sent.
	Timestamp time.Time
}

// Communication is a struct that contains the communication channels for the agent.
type Communication struct {
	Commands chan toolmanager.Command
	Messages chan Message
	Status   chan Status
}

// RunOptions is a struct that contains the options for the Run method.
type RunOptions struct {
	// Task is the task to be executed.
	Task string
	// Prompt is an optional prompt to be used for the agent instead of the default one.
	Prompt string
	// Caller is an optional tool that is calling the agent.
	Caller string
	// Tools is an optional list of tools to be used by the agent.
	Tools map[string]toolmanager.Tool
}

// Option is a function that configures the Agent.
type Option func(*Agent)

const (
	// Name is the name of the agent.
	Name = "Sredo"
)

// New creates a new Agent.
func New(opts ...Option) *Agent {
	a := &Agent{
		ctx:    context.Background(),
		cfg:    config.New().GetConfig(),
		logger: slog.New(slog.DiscardHandler),
		communication: &Communication{
			Commands: make(chan toolmanager.Command),
			Messages: make(chan Message),
			Status:   make(chan Status),
		},
	}

	for _, opt := range opts {
		opt(a)
	}

	if a.client == nil && a.cfg.Anthropic.APIKey != "" {
		a.client = anthropic.NewClient(option.WithAPIKey(a.cfg.Anthropic.APIKey))
	}

	a.logger.WithGroup("config").With("max_tokens", a.cfg.Anthropic.MaxTokens).With("model", a.cfg.Anthropic.Model).
		With("temperature", a.cfg.Anthropic.Temperature).Debug("Agent initialized.")

	return a
}

// WithContext sets the context for the agent.
func WithContext(ctx context.Context) Option {
	return func(a *Agent) {
		a.ctx = ctx
	}
}

// WithConfig sets the configuration for the agent.
func WithConfig(cfg config.Configuration) Option {
	return func(a *Agent) {
		a.cfg = cfg
	}
}

// WithLogger sets the logger for the agent.
func WithLogger(logger *slog.Logger) Option {
	return func(a *Agent) {
		a.logger = logger.With("component", "agent")
	}
}

// WithClient sets the client for the agent.
func WithClient(client *anthropic.Client) Option {
	return func(a *Agent) {
		a.client = client
	}
}

// WithCommunication sets the communication channels for the agent.
func WithCommunication(communication *Communication) Option {
	return func(a *Agent) {
		a.communication = communication
	}
}

// Run runs the agent with the given task and tools.
func (a *Agent) Run(opts *RunOptions, ctx context.Context) ([]toolmanager.ToolOutput, error) {
	if opts == nil {
		return nil, fmt.Errorf(ErrNoRunOptions)
	}

	if opts.Task == "" {
		return nil, fmt.Errorf(ErrNoTaskProvided)
	}

	if ctx == nil {
		ctx = a.ctx
	}

	prompt := agentPrompt
	if opts.Prompt != "" {
		prompt = opts.Prompt
	}

	logger := a.logger.With("task", opts.Task).With("tool", opts.Caller).With("tools.count", len(opts.Tools))
	logger.Debug("Agent running.")
	a.communication.Status <- StatusRunning

	output := []toolmanager.ToolOutput{}
	messages := []anthropic.MessageParam{anthropic.NewUserMessage(anthropic.NewTextBlock(opts.Task))}

	for {
		msg := anthropic.MessageNewParams{
			Model:     anthropic.F(a.cfg.Anthropic.Model),
			MaxTokens: anthropic.F(a.cfg.Anthropic.MaxTokens),
			System: anthropic.F([]anthropic.TextBlockParam{
				anthropic.NewTextBlock(prompt),
			}),
			Messages:    anthropic.F(messages),
			Tools:       anthropic.F(convertTools(opts.Tools)),
			Temperature: anthropic.F(a.cfg.Anthropic.Temperature),
		}

		if len(opts.Tools) > 0 {
			msg.ToolChoice = anthropic.F(anthropic.ToolChoiceUnionParam(anthropic.ToolChoiceAutoParam{
				DisableParallelToolUse: anthropic.F(true),
				Type:                   anthropic.F(anthropic.ToolChoiceAutoTypeAuto),
			}))
		}

		message, err := a.client.Messages.New(ctx, msg)

		if err != nil {
			logger.With("error", err).Error("Failed to send message to Anthropic API.")
			continue
		}

		toolResults := []anthropic.ContentBlockParamUnion{}
		for _, block := range message.Content {
			switch block := block.AsUnion().(type) {
			case anthropic.TextBlock:
				if a.communication.Messages != nil {
					a.communication.Messages <- Message{
						Tool:      opts.Caller,
						Message:   block.Text,
						Timestamp: time.Now(),
					}
					// TODO(t-dabasinskas): Remove this once we update UI
					logger.With("message", block.Text).Debug("Agent message.")
				}
			case anthropic.ToolUseBlock:
				isError := false
				toolInputs := map[string]any{}

				if err := json.Unmarshal(block.Input, &toolInputs); err != nil {
					logger.With("error", err).Error("Failed to unmarshal tool inputs.")
					continue
				}

				var toolOutput *toolmanager.ToolOutput
				tool, ok := opts.Tools[block.Name]
				if !ok {
					logger.With("tool_name", block.Name).Warn("Tool not found, skipping.")
					continue
				}

				toolOutput, err = tool.Execute(toolInputs, ctx)
				if err != nil {
					logger.With("error", err).Error("Failed to execute tool.")
					isError = true
				}

				if toolOutput == nil {
					logger.With("tool_name", block.Name).Warn("Tool has no output, skipping.")
					continue
				}

				output = append(output, *toolOutput)

				if toolOutput.Result != "" {
					a.communication.Messages <- Message{
						Tool:      opts.Caller,
						Message:   toolOutput.Result,
						Timestamp: time.Now(),
					}
					// TODO(t-dabasinskas): Remove this once we update UI
					logger.With("tool_message", toolOutput.Result).Debug("Tool message.")
				}

				if toolOutput.ExecutedCommand != nil {
					a.communication.Commands <- *toolOutput.ExecutedCommand
					// TODO(t-dabasinskas): Remove this once we update UI
					logger.With("command", toolOutput.ExecutedCommand).Debug("Tool command.")
				}

				outputJSON, err := json.Marshal(toolOutput)
				if err != nil {
					logger.With("error", err).Error("Failed to marshal tool output.")
					continue
				}

				resultBlock := anthropic.NewToolResultBlock(block.ID, string(outputJSON), isError)
				toolResults = append(toolResults, resultBlock)
			}
		}

		messages = append(messages, message.ToParam())
		if len(toolResults) == 0 {
			break
		}

		messages = append(messages, anthropic.NewUserMessage(toolResults...))
	}

	return output, nil
}

// convertTools converts the tools to the format required by the Anthropic SDK.
func convertTools(tools map[string]toolmanager.Tool) (anthropicTools []anthropic.ToolParam) {
	for _, t := range tools {
		anthropicTools = append(anthropicTools, anthropic.ToolParam{
			Name:        anthropic.F(t.GetName()),
			Description: anthropic.F(t.GetDescription()),
			InputSchema: anthropic.F(any(t.GetInputSchema())),
		})
	}

	return
}
