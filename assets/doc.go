// Package assets provides embedded static assets for the sredo application.
//
// The package uses Go's embed functionality to include various static assets
// that are required for the application to function. These assets are compiled
// into the binary, ensuring they are always available at runtime.
//
// # Embedded Assets
//
// The package contains two main categories of embedded assets:
//
// Themes Directory (/themes):
//   - Contains theme configuration files in YAML format
//   - Includes default.yaml which defines the default application theme
//   - Themes are used to customize the appearance of the terminal UI
//
// Tools Directory (/tools):
//   - Contains tool-specific configuration files in YAML format
//   - Includes git.yaml which defines Git-related configurations and commands
//   - Tools configurations define how sredo interacts with various development tools
//
// # Usage
//
// The assets are exposed through two embedded filesystems:
//
//	var Themes embed.FS // Access to theme configurations
//	var Tools embed.FS  // Access to tool configurations
//
// To access these assets in other parts of the application, import this package
// and use the appropriate embedded filesystem variable. The files can be read
// using standard fs.FS operations.
//
// Example:
//
//	themeData, err := assets.Themes.ReadFile("themes/default.yaml")
//	toolData, err := assets.Tools.ReadFile("tools/git.yaml")
package assets
