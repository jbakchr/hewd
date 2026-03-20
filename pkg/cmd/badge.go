package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
documentation. Colors are automatically chosen based on the score to provide a
quick visual indication of repository health.

Badge generation is fully local—no external services are required. This makes
it ideal for README badges, CI pipelines, or publishing artifacts in dashboards.
The SVG can be written to any path (e.g., docs/badge.svg or .github/badges/).`,
		Example: `
  # Generate an SVG badge for the current project
  hewd badge --output badge.svg

  # Write the badge into a documentation folder
  hewd badge --output docs/health-badge.svg

  # Regenerate the badge after running diagnostics
  hewd doctor --json > report.json
  hewd badge --output badge.svg

  # Use badge generation inside CI
  hewd badge --output badge.svg
`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if output == "" {
				return fmt.Errorf("--output is required (example: --output badge.svg)")
			}

			// Ensure output directory exists
			outDir := filepath.Dir(output)
			if outDir != "." && outDir != "" {
				if err := os.MkdirAll(outDir, 0755); err != nil {
					return fmt.Errorf("failed to create output directory %s: %w", outDir, err)
				}
			}

			// Determine working directory
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not determine working directory: %w", err)
			}

			// Load config (optional)
			cfg, _ := config.Load(cwd)

			// Scan repository
			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			// Run full rule engine
			results := rules.RunAll(summary, cfg, nil, nil)

			// Compute score
			scored := score.Score(results, cfg)

			// Generate badge SVG
			svg := badge.Generate(scored)

			// Write file
			if err := os.WriteFile(output, []byte(svg), 0644); err != nil {
				return fmt.Errorf("failed to write badge to %s: %w", output, err)
			}

			fmt.Printf("Badge generated → %s\n", output)
			return nil
		},
	}

	// Flags
	cmd.Flags().StringVar(&output, "output", "", "Output path for the generated SVG badge (required)")

	return cmd
}
