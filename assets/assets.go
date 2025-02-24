package assets

import (
	"bytes"
	"embed"
	"html/template"
)

var (
	//go:embed themes
	Themes embed.FS
	//go:embed tools
	Tools embed.FS

	// ToolsDir is the directory containing the tools.
	ToolsDir = "tools"
	// ThemeDir is the directory containing the themes.
	ThemeDir = "themes"

	//go:embed prompts/agent_system.tmpl
	agentSystemPrompt string
	//go:embed prompts/tool_system.tmpl
	toolSystemPrompt string
	//go:embed prompts/tool_user.tmpl
	toolUserPrompt string
)

const (
	// ErrToolRenderingPrompt is the error returned when a prompt cannot be rendered.
	ErrToolRenderingPrompt = "prompt cannot be rendered"
)

// AgentSystemPromptData is the data for the agent system prompt.
type AgentSystemPromptData struct {
	// Shell is the shell to use for the agent.
	Shell string
}

// ToolSystemPromptData is the data for the tool system prompt.
type ToolSystemPromptData struct {
	// Shell is the shell to use for the tool.
	Shell string
	// Name is the name of the tool.
	Name string
	// Executable is the executable to use for the tool.
	Executable string
	// Rules are the rules for the tool.
	Rules []string
}

// ToolUserPromptData is the data for the tool user prompt.
type ToolUserPromptData struct {
	// Task is the task to complete.
	Task string
	// Params are the parameters for the tool.
	Params map[string]any
	// Context is the context for the tool.
	Context map[string]string
	// WorkingDirectory is the working directory for the tool.
	WorkingDirectory string
}

// RenderAgentSystemPrompt renders the agent system prompt.
func RenderAgentSystemPrompt(data *AgentSystemPromptData) (string, error) {
	return render("agent_system", agentSystemPrompt, data)
}

// RenderToolSystemPrompt renders the tool system prompt.
func RenderToolSystemPrompt(data *ToolSystemPromptData) (string, error) {
	return render("tool_system", toolSystemPrompt, data)
}

// RenderToolUserPrompt renders the tool user prompt.
func RenderToolUserPrompt(data *ToolUserPromptData) (string, error) {
	return render("tool_user", toolUserPrompt, data)
}

// render is a generic function that renders a template with the given data.
func render(templateName, templateContent string, data any) (string, error) {
	tmpl, err := template.New(templateName).Parse(templateContent)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
