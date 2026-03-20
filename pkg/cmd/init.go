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
		Short: "Initialize a new hewd configuration in the current repository.",
		Long: `hewd init creates a new .hewd/config.yaml file in the current repository.
The configuration file allows you to customize rule behavior, severity levels,
category filtering, and scan include/exclude paths used by all hewd commands.

Running 'hewd init' is optional—hewd works without configuration—but creating
a config file is recommended for teams that want consistent project standards,
CI behavior, or customized scoring rules tailored to their repositories.

The command is safe to run multiple times. It will not overwrite an existing 
configuration unless the --force flag is explicitly provided.`,
		Example: `
  # Initialize hewd configuration (recommended for new repos)
  hewd init

  # Initialize config with explicit confirmation
  hewd init --force

  # View generated config
  cat .hewd/config.yaml

  # Modify rules or severity levels after initialization
  nano .hewd/config.yaml
`,
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
