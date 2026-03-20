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
		Short: "Generate an SVG badge displaying the project's overall health score.",
		Long: `hewd badge generates a standalone SVG badge representing the project's
overall health score as calculated by 'hewd doctor'. The badge is similar in
style to common README badges and can be embedded directly into Markdown or
documentation. Colors are automatically chosen based on the score to indicate
project health at a glance.

Badge generation does not require any external services; everything is rendered
locally and written to a file. This is ideal for adding a project health badge
to your README, publishing badges in CI pipelines, or generating artifacts for
dashboards.`,
		Example: `
  # Generate an SVG badge for the current project
  hewd badge --output badge.svg

  # Write the badge to a docs folder
  hewd badge --output docs/health-badge.svg

  # Regenerate the badge after running diagnostics
  hewd doctor --json > report.json
  hewd badge --output badge.svg

  # Use badge generation inside CI
  hewd badge --output badge.svg
`,
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

	cmd.Flags().StringVar(&output, "output", "", "Output path for the generated SVG badge (e.g., badge.svg)")

	return cmd
}
