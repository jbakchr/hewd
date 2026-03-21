package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/jbakchr/hewd/internal/badge"
	"github.com/jbakchr/hewd/internal/cliutils"
	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/helptext"
	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/scan"
	"github.com/jbakchr/hewd/internal/score"
)

func newBadgeCmd() *cobra.Command {
	var output string

	cmd := &cobra.Command{
		Use:     helptext.BadgeUse,
		Short:   helptext.BadgeShort,
		Long:    helptext.BadgeLong,
		Example: helptext.BadgeExample,

		RunE: func(cmd *cobra.Command, args []string) error {

			// require an output file
			if output == "" {
				return cliutils.ErrHint(
					"--output is required",
					"example: hewd badge --output badge.svg",
				)
			}

			// ensure the parent directory exists
			dir := filepath.Dir(output)
			if dir != "." {
				if err := os.MkdirAll(dir, 0755); err != nil {
					return cliutils.ErrHint(
						fmt.Sprintf("failed to create directory %s: %v", dir, err),
						"ensure you have write permissions for the target path",
					)
				}
			}

			// determine working directory
			cwd, err := os.Getwd()
			if err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to determine working directory: %v", err),
					"ensure the current directory is accessible",
				)
			}

			// load config
			cfg, err := config.Load(cwd)
			if err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to load configuration: %v", err),
					"ensure .hewd/config.yaml is valid yaml and readable",
				)
			}

			// scan repository
			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to scan directory: %v", err),
					"ensure you are running hewd inside a valid repository",
				)
			}

			// run all rules
			results := rules.RunAll(summary, cfg, nil, nil)

			// compute score
			overallScore := score.Score(results, cfg)

			// generate badge SVG
			svg := badge.Generate(overallScore)

			// write to file
			if err := os.WriteFile(output, []byte(svg), 0644); err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to write file %s: %v", output, err),
					"ensure directory permissions allow writing",
				)
			}

			fmt.Printf("Badge generated → %s\n", output)
			return nil
		},
	}

	cmd.GroupID = "reporting"

	cmd.Flags().StringVar(
		&output,
		"output",
		"",
		"Write the generated SVG badge to the specified file path (required).",
	)

	return cmd
}
