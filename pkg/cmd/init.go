package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/jbakchr/hewd/internal/cliutils"
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

			// ensure directory exists
			if err := os.MkdirAll(cfgDir, 0755); err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to create directory %s: %v", cfgDir, err),
					"ensure you have write permissions in the project root",
				)
			}

			// warn if config exists unless forced
			if _, err := os.Stat(cfgPath); err == nil && !force {
				return cliutils.ErrHint(
					".hewd/config.yaml already exists",
					"use --force to overwrite the existing file",
				)
			}

			// default configuration template
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

			// write config
			if err := os.WriteFile(cfgPath, []byte(defaultCfg), 0644); err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to write file %s: %v", cfgPath, err),
					"ensure you have write permissions in this directory",
				)
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
