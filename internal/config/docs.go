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
//	}
//
// Usage:
//
//	cfg := config.New()
//	if err := cfg.LoadConfig(); err != nil {
//	    log.Fatal(err)
//	}
//	config := cfg.GetConfig()
//
// Environment Variables:
//   - ANTHROPIC_API_KEY: API key for Anthropic
//   - SREDO_UI_THEME: UI theme name
//   - SREDO_LOGGING_LEVEL: Log level (debug, info, warn, error)
//   - SREDO_ANTHROPIC_MODEL: Model name
//   - SREDO_ANTHROPIC_TEMPERATURE: Temperature value
//
// Directory Structure:
//
//	~/.sredo/
//	├── config.yaml  // Configuration file
//	├── log.log     // Default log file
//	├── cache/      // Cache directory
//	├── themes/     // Custom UI themes
//	└── tools/      // Tool-specific data
package config
