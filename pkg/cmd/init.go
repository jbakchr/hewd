package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize hewd configuration",
		Long: `Create the default hewd configuration directory
and bootstrap initial configuration files.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			configDir, err := os.UserConfigDir()
			if err != nil {
				return fmt.Errorf("could not determine config directory: %w", err)
			}

			hewdDir := filepath.Join(configDir, "hewd")
			if err := os.MkdirAll(hewdDir, 0o755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", hewdDir, err)
			}

			configFile := filepath.Join(hewdDir, "config.yaml")

			if _, err := os.Stat(configFile); os.IsNotExist(err) {
				contents := []byte(`# hewd configuration file
# Add your settings below.

`)
				if err := os.WriteFile(configFile, contents, 0o644); err != nil {
					return fmt.Errorf("failed to write config file: %w", err)
				}

				fmt.Println("Initialized hewd configuration at:")
				fmt.Println(" ", configFile)
			} else {
				fmt.Println("hewd configuration already exists at:")
				fmt.Println(" ", configFile)
			}

			return nil
		},
	}
}
