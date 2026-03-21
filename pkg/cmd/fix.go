package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jbakchr/hewd/internal/cliutils"
	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/fix"
	"github.com/jbakchr/hewd/internal/helptext"
	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/scan"
)

func newFixCmd() *cobra.Command {
	var apply bool

	cmd := &cobra.Command{
		Use:     helptext.FixUse,
		Short:   helptext.FixShort,
		Long:    helptext.FixLong,
		Example: helptext.FixExample,

		RunE: func(cmd *cobra.Command, args []string) error {

			// Determine working directory
			cwd, err := os.Getwd()
			if err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to determine working directory: %v", err),
					"ensure the current directory is accessible",
				)
			}

			// Load config (optional)
			cfg, err := config.Load(cwd)
			if err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to load configuration: %v", err),
					"ensure .hewd/config.yaml is valid yaml and readable",
				)
			}

			// Scan project
			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to scan directory: %v", err),
					"ensure you are running hewd inside a valid repository",
				)
			}

			// Run all rules — fix always considers all categories
			results := rules.RunAll(summary, cfg, nil, nil)

			// Detect fixable items
			fixables := fix.DetectFixes(results, cwd)
			if len(fixables) == 0 {
				return cliutils.Err("no fixable issues found")
			}

			// Dry‑run mode (default)
			if !apply {
				fmt.Println("Fixable issues:")
				for _, f := range fixables {
					fmt.Printf("  - %s → %s\n", f.RuleID, f.Message)
				}
				fmt.Println("\nRun with --apply to write these changes to disk.")
				return nil
			}

			// Apply fixes to disk
			for _, f := range fixables {
				fmt.Printf("Applying fix for %s...\n", f.RuleID)
				if err := fix.ApplyFix(f); err != nil {
					return cliutils.ErrHint(
						fmt.Sprintf("failed to apply fix for %s: %v", f.RuleID, err),
						"verify file permissions and repository structure",
					)
				}
			}

			fmt.Println("All fixes applied.")
			return nil
		},
	}

	cmd.GroupID = "maintenance"

	cmd.Flags().BoolVar(&apply, "apply", false, "Apply fixes and write new files to disk (default is dry-run).")

	return cmd
}
