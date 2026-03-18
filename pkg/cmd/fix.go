package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/fix"
	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/scan"
)

func newFixCmd() *cobra.Command {
	var apply bool

	cmd := &cobra.Command{
		Use:   "fix",
		Short: "Automatically fix common documentation and configuration issues",
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, _ := os.Getwd()
			cfg, _ := config.Load(cwd)

			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			// No filtering of categories here = full run
			results := rules.RunAll(summary, cfg, nil, nil)

			fixes := fix.DetectFixes(results, cwd)

			if len(fixes) == 0 {
				fmt.Println("No fixable issues found.")
				return nil
			}

			if !apply {
				fmt.Println("Fixable issues:")
				for _, f := range fixes {
					fmt.Printf("  - %s → %s\n", f.RuleID, f.Message)
				}
				fmt.Println("\nRun with --apply to apply changes.")
				return nil
			}

			for _, f := range fixes {
				fmt.Printf("Applying fix for %s...\n", f.RuleID)
				if err := fix.ApplyFix(f); err != nil {
					return err
				}
			}

			fmt.Println("All fixes applied.")
			return nil
		},
	}

	cmd.Flags().BoolVar(&apply, "apply", false, "Apply fixes instead of dry-run")

	return cmd
}
