package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jbakchr/hewd/internal/config"
)

// loadedConfig is cached so that multiple commands (scan, doctor)
// do not reload the config file independently.
var loadedConfig *config.Config

// LoadHewdConfig loads `.hewd/config.yaml` once and reuses it.
func LoadHewdConfig() *config.Config {
	if loadedConfig != nil {
		return loadedConfig
	}

	// Attempt to load config; missing config is not an error
	cfg, err := config.Load(".")
	if err == nil {
		loadedConfig = cfg
	}
	return loadedConfig
}

// NewRootCmd constructs the top-level Cobra command for hewd.
func NewRootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "hewd",
		Short: "hewd — Documentation & Project Structure Toolkit",
		Long: `hewd is a CLI tool that scans your project for documentation,
configuration, and structural assets, and performs diagnostics to help maintain
healthy codebases. Configuration is loaded from .hewd/config.yaml when present.`,
	}

	// Attach commands
	rootCmd.AddCommand(newScanCmd())
	rootCmd.AddCommand(newDoctorCmd())
	rootCmd.AddCommand(newInitCmd())

	// Version command
	rootCmd.Version = version
	rootCmd.SetVersionTemplate("hewd version {{.Version}}\n")

	return rootCmd
}

// Execute is the main entry point invoked from main.go.
func Execute(version string) {
	if err := NewRootCmd(version).Execute(); err != nil {
		fmt.Println(err)
	}
}
