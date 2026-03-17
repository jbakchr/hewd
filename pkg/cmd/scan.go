package cmd

import (
	"fmt"
	"os"

	iscan "github.com/jbakchr/hewd/internal/scan"

	"github.com/spf13/cobra"
)

//
// ───────────────────────────────────────────────────────────────
//   Cobra command setup
// ───────────────────────────────────────────────────────────────
//

func newScanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Scan the current directory and summarize the project",
		Long: `Scan the current working directory and produce a project summary,
including detected languages, documentation files, file counts, 
configuration files, and structure info.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not get working directory: %w", err)
			}

			summary, err := iscan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			jsonOut, _ := cmd.Flags().GetBool("json")
			yamlOut, _ := cmd.Flags().GetBool("yaml")
			pretty, _ := cmd.Flags().GetBool("pretty")

			if jsonOut && yamlOut {
				return fmt.Errorf("cannot use --json and --yaml together")
			}

			if jsonOut {
				return printJSON(summary, pretty)
			}
			if yamlOut {
				return printYAML(summary)
			}

			printScanSummary(summary)
			return nil
		},
	}

	cmd.Flags().Bool("json", false, "Output scan result in JSON format")
	cmd.Flags().Bool("yaml", false, "Output scan result in YAML format")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON output")

	return cmd
}
