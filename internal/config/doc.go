// Package config provides configuration management for the sredo CLI application.
//
// The package handles:
//   - Loading configuration from YAML files
//   - Environment variable binding
//   - Configuration validation
//   - Default values
//   - Directory structure setup
//   - Logging setup
//
// Configuration Structure:
//
//	Configuration {
//	  UI:        UIConfiguration        // UI theme and styling
//	  Logging:   LoggingConfiguration   // Log file path and level
//	  Anthropic: AnthropicConfiguration // API settings for Anthropic
//	  Tools:     ToolsConfiguration     // Global tool settings
//	}
//
// Usage:
//
//	manager := config.New()
//	if err := manager.LoadConfig(); err != nil {
//	    log.Fatal(err)
//	}
//	config := manager.GetConfig()
//
// Environment Variables:
//   - ANTHROPIC_API_KEY: API key for Anthropic
//   - SREDO_UI_THEME: UI theme name
//   - SREDO_LOGGING_LEVEL: Log level (debug, info, warn, error)
//   - SREDO_ANTHROPIC_MODEL: Model name
//   - SREDO_ANTHROPIC_TEMPERATURE: Temperature value
//   - SREDO_ANTHROPIC_MAX_TOKENS: Maximum tokens for completion
//   - SREDO_TOOLS_TIMEOUT: Global timeout for tools in seconds
//   - SREDO_TOOLS_EXEC_TIMEOUT: Timeout for exec tool in seconds
//
// Directory Structure:
//
//	~/.sredo/
//	├── config.yaml  // Configuration file
//	├── log.log     // Default log file
//	├── cache/      // Cache directory
//	└── tools/      // Tool-specific data
//
// The package uses the following error constants for error handling:
//   - ErrCreateDirs: Returned when directory creation fails
//   - ErrReadConfig: Returned when config file cannot be read
//   - ErrUnmarshalConfig: Returned when config parsing fails
//   - ErrMissingAPIKey: Returned when Anthropic API key is missing
//   - ErrInvalidTemp: Returned when temperature is not between 0 and 1
//   - ErrInvalidMaxTokens: Returned when max tokens is not positive
//   - ErrInvalidLogLevel: Returned when log level is invalid
//   - ErrOpenLogFile: Returned when log file cannot be opened
package config
