package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func newInitCmd() *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new hewd configuration in the current repository.",
		Long: `hewd init creates a .hewd/config.yaml file in the current repository.
The configuration file allows customization of rule behavior, severity levels,
category filtering, scoring weights, and include/exclude paths used by all
hewd commands.

Running 'hewd init' is optional—hewd works without configuration—but creating
a config file is recommended for teams that want consistent project standards,
CI behavior, or customized scoring rules.

This command is safe to run multiple times. It will not overwrite an existing
configuration unless the --force flag is explicitly provided.`,
		Example: `
  # Initialize hewd configuration (recommended for new repos)
  hewd init

  # Force re-creation of the config (overwrites existing file)
  hewd init --force

  # View generated config
  cat .hewd/config.yaml

  # Modify rules or severity levels after initialization
  nano .hewd/config.yaml
`,
		RunE: func(cmd *cobra.Command, args []string) error {

			cfgDir := filepath.Join(".", ".hewd")
			cfgPath := filepath.Join(cfgDir, "config.yaml")

			// Ensure .hewd directory exists
			if _, err := os.Stat(cfgDir); os.IsNotExist(err) {
				if err := os.MkdirAll(cfgDir, 0755); err != nil {
					return fmt.Errorf("failed to create .hewd directory: %w", err)
				}
				fmt.Println("Created .hewd/ directory")
			}

			// Check for existing config file
			if _, err := os.Stat(cfgPath); err == nil && !force {
				fmt.Println(".hewd/config.yaml already exists — use --force to overwrite.")
				return nil
			}

			// Default template
			defaultCfg := `# hewd configuration file
#
# This file defines rule behavior, severity overrides, scoring weights, and
# include/exclude paths for scans and diagnostics. All fields are optional.

rules: {}
weights: {}

scan:
  include: []
  exclude:
    - node_modules
    - vendor
`

			// Write file (with overwrite optional)
			if err := os.WriteFile(cfgPath, []byte(defaultCfg), 0644); err != nil {
				return fmt.Errorf("failed to write config.yaml: %w", err)
			}

			if force {
				fmt.Println("Overwrote existing .hewd/config.yaml")
			} else {
				fmt.Println("Created .hewd/config.yaml")
			}

			return nil
		},
	}

	// Flags
	cmd.Flags().BoolVar(&force, "force", false,
		"Overwrite existing .hewd/config.yaml if it already exists")

	return cmd
}
