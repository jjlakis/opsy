package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Configuration is the configuration for the sredo CLI.
type Configuration struct {
	// UI is the configuration for the UI.
	UI UIConfiguration `yaml:"ui"`
	// Logging is the configuration for the logging.
	Logging LoggingConfiguration `yaml:"logging"`
	// Anthropic is the configuration for the Anthropic API.
	Anthropic AnthropicConfiguration `yaml:"anthropic"`
}

// UIConfiguration is the configuration for the UI.
type UIConfiguration struct {
	// Theme is the theme for the UI.
	Theme string `yaml:"theme"`
}

// LoggingConfiguration is the configuration for the logging.
type LoggingConfiguration struct {
	// Path is the path to the log file.
	Path string `yaml:"path"`
	// Level is the logging level.
	Level string `yaml:"level"`
}

// AnthropicConfiguration is the configuration for the Anthropic API.
type AnthropicConfiguration struct {
	// APIKey is the API key for the Anthropic API.
	APIKey string `mapstructure:"api_key" yaml:"api_key"`
	// Model is the model to use for the Anthropic API.
	Model string `yaml:"model"`
	// Temperature is the temperature to use for the Anthropic API.
	Temperature float64 `yaml:"temperature"`
	// MaxTokens is the maximum number of tokens to use for the Anthropic API.
	MaxTokens int64 `mapstructure:"max_tokens" yaml:"max_tokens"`
}

// Configurer is an interface for managing configuration.
type Configurer interface {
	// LoadConfig loads the configuration from the config file.
	LoadConfig() error
	// GetConfig returns the current configuration.
	GetConfig() Configuration
	// GetLogger returns the default logger.
	GetLogger() (*slog.Logger, error)
}

// ConfigManager is the configuration manager for the sredo CLI.
type Config struct {
	configuration Configuration
	homePath      string
}

const (
	dirConfig  = ".sredo"
	dirCache   = ".sredo/cache"
	envPrefix  = "SREDO"
	configFile = "config"
	configType = "yaml"
)

var (
	// ErrCreateConfigDir is returned when the config directory cannot be created.
	ErrCreateConfigDir = errors.New("failed to create config directory")
	// ErrCreateCacheDir is returned when the cache directory cannot be created.
	ErrCreateCacheDir = errors.New("failed to create cache directory")
	// ErrCreateDirs is returned when the directories cannot be created.
	ErrCreateDirs = errors.New("failed to create directories")
	// ErrReadConfig is returned when the config file cannot be read.
	ErrReadConfig = errors.New("failed to read config")
	// ErrUnmarshalConfig is returned when the config file cannot be unmarshalled.
	ErrUnmarshalConfig = errors.New("failed to unmarshal config")
	// ErrMissingAPIKey is returned when the Anthropic API key is missing.
	ErrMissingAPIKey = errors.New("anthropic API key is required")
	// ErrInvalidTemp is returned when the Anthropic temperature is invalid.
	ErrInvalidTemp = errors.New("anthropic temperature must be between 0 and 1")
	// ErrInvalidMaxTokens is returned when the Anthropic max tokens are invalid.
	ErrInvalidMaxTokens = errors.New("anthropic max tokens must be greater than 0")
	// ErrInvalidLogLevel is returned when the logging level is invalid.
	ErrInvalidLogLevel = errors.New("invalid logging level")
	// ErrInvalidTheme is returned when the theme is invalid.
	ErrInvalidTheme = errors.New("invalid theme")
	// ErrOpenLogFile is returned when the log file cannot be opened.
	ErrOpenLogFile = errors.New("failed to open log file")
	// ErrWriteConfig is returned when the config file cannot be written.
	ErrWriteConfig = errors.New("failed to write config")
)

// New creates a new config instance.
func New() *Config {
	homeDir, _ := os.UserHomeDir()

	config := &Config{
		homePath: homeDir,
	}

	config.setDefaults()

	viper.AutomaticEnv()
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath(filepath.Join(homeDir, dirConfig))
	viper.SetConfigName(configFile)
	viper.SetConfigType(configType)

	_ = viper.BindEnv("anthropic.api_key", "ANTHROPIC_API_KEY")

	return config
}

// LoadConfig loads the configuration from the config file.
func (c *Config) LoadConfig() error {
	if err := c.createDirs(); err != nil {
		return fmt.Errorf("%w: %v", ErrCreateDirs, err)
	}

	if err := viper.SafeWriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); !ok {
			return fmt.Errorf("%w: %v", ErrWriteConfig, err)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("%w: %v", ErrReadConfig, err)
	}

	if err := viper.Unmarshal(&c.configuration); err != nil {
		return fmt.Errorf("%w: %v", ErrUnmarshalConfig, err)
	}

	return c.validate()
}

// GetConfig returns the current configuration.
func (c *Config) GetConfig() Configuration {
	return c.configuration
}

// GetLogger returns a logger that writes to the log file.
func (c *Config) GetLogger() (*slog.Logger, error) {
	logFile, err := os.OpenFile(c.configuration.Logging.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrOpenLogFile, err)
	}

	var lvl slog.Level
	switch c.configuration.Logging.Level {
	case "debug":
		lvl = slog.LevelDebug
	case "warn":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}

	logger := slog.New(slog.NewTextHandler(logFile, &slog.HandlerOptions{
		Level: lvl,
	}))

	return logger, nil
}

func (c *Config) createDirs() error {
	if err := os.MkdirAll(filepath.Join(c.homePath, dirConfig), 0755); err != nil {
		return fmt.Errorf("%w: %v", ErrCreateConfigDir, err)
	}

	if err := os.MkdirAll(filepath.Join(c.homePath, dirCache), 0755); err != nil {
		return fmt.Errorf("%w: %v", ErrCreateCacheDir, err)
	}

	return nil
}

func (c *Config) validate() error {
	if c.configuration.Anthropic.APIKey == "" {
		return ErrMissingAPIKey
	}

	if c.configuration.Anthropic.Temperature < 0 || c.configuration.Anthropic.Temperature > 1 {
		return ErrInvalidTemp
	}

	if c.configuration.Anthropic.MaxTokens < 1 {
		return ErrInvalidMaxTokens
	}

	level := strings.ToLower(c.configuration.Logging.Level)
	validLevels := map[string]bool{
		"debug": true,
		"info":  true,
		"warn":  true,
		"error": true,
	}
	if !validLevels[level] {
		return ErrInvalidLogLevel
	}

	return nil
}

func (c *Config) setDefaults() {
	viper.SetDefault("ui.theme", "default")
	viper.SetDefault("logging.path", filepath.Join(c.homePath, dirConfig, "log.log"))
	viper.SetDefault("logging.level", "info")
	viper.SetDefault("anthropic.model", "claude-3-5-sonnet-latest")
	viper.SetDefault("anthropic.temperature", 0.9)
	viper.SetDefault("anthropic.max_tokens", 1024)
}
