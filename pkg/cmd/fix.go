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
		Short: "Automatically generate missing documentation, structure, and CI files.",
		Long: `hewd fix analyzes the repository for missing documentation, structure,
and configuration files, and automatically generates recommended assets. 
This includes files such as LICENSE, CONTRIBUTING.md, CHANGELOG.md, default CI 
workflows, and the docs/ directory.

By default, fixes are shown as a dry-run without making changes. Use the 
--apply flag to write the generated files to disk. hewd fix is safe to run 
repeatedly—existing files are never overwritten.

This command is ideal for preparing repositories for publication, enforcing 
documentation standards, and quickly bootstrapping missing project components.`,
		Example: `
  # Show which fixes would be applied (dry-run)
  hewd fix

  # Apply fixes and write files to disk
  hewd fix --apply

  # Fix missing documentation before running diagnostics
  hewd fix --apply && hewd doctor

  # Preview fixes and save output to a file
  hewd fix > fix-preview.txt
`,
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

	cmd.Flags().BoolVar(&apply, "apply", false, "Apply fixes and write new files to disk (default: dry-run)")

	return cmd
}
