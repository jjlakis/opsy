// Package thememanager provides functionality for managing and loading color themes
// for terminal user interfaces.
//
// The package supports loading themes from both embedded files and custom directories.
// Themes are defined in YAML format and contain base and accent colors that can be
// used for consistent styling across the application.
//
// Basic usage:
//
//	// Create a new theme manager with default settings
//	tm := thememanager.New()
//
//	// Load the default theme
//	err := tm.LoadTheme("")
//
//	// Or load a specific theme
//	err := tm.LoadTheme("dark")
//
//	// Get the current theme for use
//	theme := tm.GetTheme()
//
// Custom theme directory:
//
//	// Create a theme manager with a custom theme directory
//	tm := thememanager.New(thememanager.WithDirectory("/path/to/themes"))
//
// Theme files should be YAML files with the .yaml extension and follow this structure:
//
//	base:
//	  base00: "#1A1B26"  # Background
//	  base01: "#24283B"  # Light Background
//	  base02: "#292E42"  # Selection Background
//	  base03: "#565F89"  # Comments, Invisibles
//	  base04: "#A9B1D6"  # Dark Foreground
//	accent:
//	  accent0: "#FF9E64" # Orange
//	  accent1: "#9ECE6A" # Green
//	  accent2: "#7AA2F7" # Blue
//
// Color Validation:
//
// All colors in the theme must:
//   - Be valid hexadecimal color codes
//   - Start with '#' character
//   - Be present for all required fields (no missing colors)
//
// Default Theme:
//
// When no theme name is provided (empty string), the manager will:
//   - Load the "default" theme from the embedded themes
//   - Use this as the fallback theme for the application
//   - Return an error if the default theme is not found or invalid
//
// The package uses the following error constants for error handling:
//   - ErrThemeNotFound: Returned when a requested theme file cannot be found
//   - ErrReadingTheme: Returned when there's an error reading the theme file
//   - ErrParsingTheme: Returned when the theme file cannot be parsed
//   - ErrDecodingTheme: Returned when theme YAML decoding fails
//   - ErrMissingColors: Returned when required colors are missing from theme
//
// The Manager interface defines the core functionality that theme managers must implement:
//   - LoadTheme: Loads a theme by name
//   - GetTheme: Returns the currently loaded theme
//
// Thread Safety:
//
// The theme manager is safe for concurrent access:
//   - LoadTheme operations are synchronized
//   - GetTheme returns a pointer to the theme, which should be treated as read-only
//   - Theme modifications should be done through LoadTheme only
package thememanager
