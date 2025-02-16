package main

import (
	"context"
	"errors"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/datolabs-io/sredo/internal/config"
	"github.com/datolabs-io/sredo/internal/thememanager"
	"github.com/datolabs-io/sredo/internal/toolmanager"
	"github.com/datolabs-io/sredo/internal/tui"
)

const (
	// ErrNoTaskProvided is the error message for no task provided.
	ErrNoTaskProvided = "no task provided"
)

// main is the entry point for the Sredo application.
func main() {
	ctx := context.Background()

	task, err := getTask()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.New()
	if err := cfg.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	logger, err := cfg.GetLogger()
	if err != nil {
		log.Fatal(err)
	}

	logger.With("component", "main").With("task", task).Info("Started Sredo")

	themeManager := thememanager.New(thememanager.WithLogger(logger.With("component", "thememanager")))
	if err := themeManager.LoadTheme(cfg.GetConfig().UI.Theme); err != nil {
		log.Fatal(err)
	}

	toolManager := toolmanager.New(
		toolmanager.WithConfig(cfg.GetConfig()),
		toolmanager.WithLogger(logger.With("component", "toolmanager")),
		toolmanager.WithContext(ctx),
	)
	if err := toolManager.LoadTools(); err != nil {
		log.Fatal(err)
	}

	tui := tui.New(
		tui.WithTheme(themeManager.GetTheme()),
		tui.WithConfig(cfg.GetConfig()),
		tui.WithTask(task),
		tui.WithToolsCount(len(toolManager.GetTools())),
	)
	p := tea.NewProgram(tui, tea.WithAltScreen(), tea.WithMouseCellMotion(), tea.WithContext(ctx))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	logger.With("component", "main").With("task", task).Info("Sredo finished")
}

// getTask returns the task from the command line arguments.
func getTask() (string, error) {
	if len(os.Args) > 1 && os.Args[1] != "" {
		return os.Args[1], nil
	}

	return "", errors.New(ErrNoTaskProvided)
}
