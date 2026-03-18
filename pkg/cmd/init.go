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
		Short: "Initialize hewd configuration",
		Long:  `Create a .hewd/config.yaml file with default settings.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			cfgDir := filepath.Join(".", ".hewd")
			if _, err := os.Stat(cfgDir); os.IsNotExist(err) {
				if err := os.MkdirAll(cfgDir, 0755); err != nil {
					return fmt.Errorf("failed to create .hewd: %w", err)
				}
				fmt.Println("Created .hewd/ directory")
			}

			cfgPath := filepath.Join(cfgDir, "config.yaml")
			if _, err := os.Stat(cfgPath); err == nil {
				fmt.Println(".hewd/config.yaml already exists — nothing to do.")
				return nil
			}

			defaultCfg := `# hewd configuration file

rules: {}
weights: {}

scan:
  include: []
  exclude:
    - node_modules
    - vendor
`

			if err := os.WriteFile(cfgPath, []byte(defaultCfg), 0644); err != nil {
				return fmt.Errorf("failed to write config.yaml: %w", err)
			}

			fmt.Println("Created .hewd/config.yaml")
			return nil
		},
	}

	return cmd
}
