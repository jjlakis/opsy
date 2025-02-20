package tool

import (
	"context"
)

// Runner is an interface that defines the methods for an agent.
type Runner interface {
	Run(opts *RunOptions, ctx context.Context) ([]Output, error)
}

// RunOptions is a struct that contains the options for runner run.
type RunOptions struct {
	// Task is the task to be executed.
	Task string
	// Prompt is an optional prompt to be used for the agent instead of the default one.
	Prompt string
	// Caller is an optional tool that is calling the agent.
	Caller string
	// Tools is an optional list of tools to be used by the agent.
	Tools map[string]Tool
}
