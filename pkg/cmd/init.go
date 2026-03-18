package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func newInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize hewd configuration for the current project",
		Long: `Initializes project-level configuration used by hewd.
Creates a .hewd/config.yaml file if it does not already exist.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			// Create .hewd directory if missing
			cfgDir := filepath.Join(".", ".hewd")
			if _, err := os.Stat(cfgDir); os.IsNotExist(err) {
				if err := os.MkdirAll(cfgDir, 0755); err != nil {
					return fmt.Errorf("failed to create .hewd directory: %w", err)
				}
				fmt.Println("Created .hewd/ directory")
			}

			// Create config.yaml only if not existing
			cfgPath := filepath.Join(cfgDir, "config.yaml")
			if _, err := os.Stat(cfgPath); err == nil {
				fmt.Println(".hewd/config.yaml already exists — nothing to do.")
				return nil
			}

			defaultConfig := `# hewd configuration file
# Adjust rule enabling, disabling, and severity overrides.
# Adjust scanning include/exclude directories.

rules: {}
weights: {}

scan:
  include: []
  exclude:
    - node_modules
    - vendor
`

			if err := os.WriteFile(cfgPath, []byte(defaultConfig), 0644); err != nil {
				return fmt.Errorf("failed to write config.yaml: %w", err)
			}

			fmt.Println("Created .hewd/config.yaml")

			return nil
		},
	}

	return cmd
}
