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
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			dir := filepath.Join(cwd, ".hewd")
			os.MkdirAll(dir, 0755)

			cfgPath := filepath.Join(dir, "config.yaml")

			if _, err := os.Stat(cfgPath); err == nil {
				return fmt.Errorf("config.yaml already exists")
			}

			defaultConfig := `rules:
  DOC-001: true
  DOC-002: true
  CFG-001: true

weights:
  DOC-001: 10
  DOC-002: 5
  CFG-001: 5

scan:
  include: []
  exclude:
    - node_modules
    - vendor
    - dist
`

			return os.WriteFile(cfgPath, []byte(defaultConfig), 0644)
		},
	}

	return cmd
}
