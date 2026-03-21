package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/jbakchr/hewd/internal/helptext"
)

func newInitCmd() *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:     helptext.InitUse,
		Short:   helptext.InitShort,
		Long:    helptext.InitLong,
		Example: helptext.InitExample,

		RunE: func(cmd *cobra.Command, args []string) error {

			cfgDir := filepath.Join(".", ".hewd")
			cfgPath := filepath.Join(cfgDir, "config.yaml")

			// Create .hewd directory if missing
			if _, err := os.Stat(cfgDir); os.IsNotExist(err) {
				if err := os.MkdirAll(cfgDir, 0755); err != nil {
					return fmt.Errorf("failed to create .hewd directory: %w", err)
				}
				fmt.Println("Created .hewd/ directory")
			}

			// If config already exists and no --force, stop
			if _, err := os.Stat(cfgPath); err == nil && !force {
				fmt.Println(".hewd/config.yaml already exists — use --force to overwrite.")
				return nil
			}

			// Default configuration template
			defaultCfg := `# hewd configuration file
#
# Defines rule behavior, severity overrides, scoring weights, and
# include/exclude paths for scans and diagnostics. All fields are optional.

rules: {}
weights: {}

scan:
  include: []
  exclude:
    - node_modules
    - vendor
`

			// Write config to disk
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

	cmd.GroupID = "maintenance"

	cmd.Flags().BoolVar(
		&force,
		"force",
		false,
		"Overwrite existing .hewd/config.yaml if it already exists.",
	)

	return cmd
}
