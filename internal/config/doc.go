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
//	  Tools:     ToolsConfiguration     // Global tool settings and exec configuration
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
//   - SREDO_TOOLS_EXEC_SHELL: Shell to use for command execution
//
// Directory Structure:
//
//	~/.sredo/
//	├── config.yaml  // Configuration file
//	├── log.log     // Default log file
//	├── cache/      // Cache directory for temporary files
//	└── tools/      // Tool-specific data and configurations
//
// The package uses the following error constants for error handling:
//   - ErrCreateDirs: Returned when directory creation fails
//   - ErrCreateConfigDir: Returned when config directory creation fails
//   - ErrCreateCacheDir: Returned when cache directory creation fails
//   - ErrReadConfig: Returned when config file cannot be read
//   - ErrWriteConfig: Returned when config file cannot be written
//   - ErrUnmarshalConfig: Returned when config parsing fails
//   - ErrValidateConfig: Returned when configuration validation fails
//   - ErrMissingAPIKey: Returned when Anthropic API key is missing
//   - ErrInvalidTemp: Returned when temperature is not between 0 and 1
//   - ErrInvalidMaxTokens: Returned when max tokens is not positive
//   - ErrInvalidLogLevel: Returned when log level is invalid
//   - ErrInvalidTheme: Returned when UI theme is invalid
//   - ErrInvalidShell: Returned when exec shell is invalid or not found
//   - ErrOpenLogFile: Returned when log file cannot be opened
//
// Validation:
//
// The package performs extensive validation of the configuration:
//   - Anthropic API key must be provided
//   - Temperature must be between 0 and 1
//   - Max tokens must be positive
//   - Log level must be one of: debug, info, warn, error
//   - UI theme must be a valid theme name
//   - Exec shell must be a valid and executable shell path
//
// Thread Safety:
//
// The configuration is safe for concurrent access after loading.
// The GetConfig method returns a copy of the configuration to prevent
// race conditions.
package config
