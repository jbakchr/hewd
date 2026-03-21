package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

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

			// Load working directory
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not determine working directory: %w", err)
			}

			cfg, _ := config.Load(cwd)

			// Scan project
			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			// Run full rule set — fix always considers all categories
			results := rules.RunAll(summary, cfg, nil, nil)

			// Detect fixable items
			fixes := fix.DetectFixes(results, cwd)

			if len(fixes) == 0 {
				fmt.Println("No fixable issues found.")
				return nil
			}

			// -------------------------
			// DRY-RUN (default)
			// -------------------------
			if !apply {
				fmt.Println("Fixable issues:")
				for _, f := range fixes {
					fmt.Printf("  - %s → %s\n", f.RuleID, f.Message)
				}
				fmt.Println("\nRun with --apply to apply these fixes.")
				return nil
			}

			// -------------------------
			// APPLY FIXES
			// -------------------------
			for _, f := range fixes {
				fmt.Printf("Applying fix for %s...\n", f.RuleID)
				if err := fix.ApplyFix(f); err != nil {
					return fmt.Errorf("failed to apply fix for %s: %w", f.RuleID, err)
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
