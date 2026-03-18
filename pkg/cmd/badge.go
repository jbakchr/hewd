package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jbakchr/hewd/internal/badge"
	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/scan"
	"github.com/jbakchr/hewd/internal/score"
)

func newBadgeCmd() *cobra.Command {
	var output string

	cmd := &cobra.Command{
		Use:   "badge",
		Short: "Generate a hewd score badge",
		RunE: func(cmd *cobra.Command, args []string) error {

			if output == "" {
				return fmt.Errorf("--output is required")
			}

			cwd, _ := os.Getwd()
			cfg, _ := config.Load(cwd)

			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			results := rules.RunAll(summary, cfg, nil, nil)
			scored := score.Score(results, cfg)

			svg := badge.Generate(scored)

			return os.WriteFile(output, []byte(svg), 0644)
		},
	}

	cmd.Flags().StringVar(&output, "output", "", "Path to write badge.svg")

	return cmd
}
