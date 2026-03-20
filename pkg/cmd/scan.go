package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/jbakchr/hewd/internal/scan"
)

func newScanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Scan the repository and detect documentation, config, languages, and structure indicators.",
		Long: `hewd scan performs a fast, lightweight analysis of the current repository.
It detects documentation files, configuration files, programming languages,
project metadata, and structural indicators such as the presence of a docs/
directory or CI workflows.

This command provides a high-level overview of your project's documentation
and structure. It is often the first step before running 'hewd doctor' or
exporting a machine-readable health report.

Scan output supports multiple formats including:

  • Pretty (human-readable)
  • JSON
  • YAML

This command is safe to run in both local development and CI environments.`,
		Example: `
  # Scan the current repository (pretty output)
  hewd scan --pretty

  # Scan and output JSON
  hewd scan --json

  # Scan and output YAML
  hewd scan --yaml

  # Save JSON output to a file
  hewd scan --json > scan.json

  # Use scan as part of a workflow before doctor
  hewd scan --pretty && hewd doctor
`,
		RunE: func(cmd *cobra.Command, args []string) error {

			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not get working directory: %w", err)
			}

			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			jsonOut, _ := cmd.Flags().GetBool("json")
			yamlOut, _ := cmd.Flags().GetBool("yaml")
			pretty, _ := cmd.Flags().GetBool("pretty")

			if jsonOut && yamlOut {
				return fmt.Errorf("cannot combine --json and --yaml")
			}

			if yamlOut && pretty {
				return fmt.Errorf("cannot combine --yaml and --pretty (pretty mode only applies to JSON)")
			}

			if jsonOut {
				var data []byte
				if pretty {
					data, err = json.MarshalIndent(summary, "", "  ")
				} else {
					data, err = json.Marshal(summary)
				}
				if err != nil {
					return err
				}
				fmt.Println(string(data))
				return nil
			}

			if yamlOut {
				data, err := yaml.Marshal(summary)
				if err != nil {
					return err
				}
				fmt.Println(string(data))
				return nil
			}

			printScanSummary(summary)
			return nil
		},
	}

	cmd.Flags().BoolP("json", "j", false, "Output results in JSON format. Use --pretty for indented output.")
	cmd.Flags().BoolP("yaml", "y", false, "Output results in YAML format.")
	cmd.Flags().BoolP("pretty", "p", false, "Pretty output for human-readable CLI usage.")

	return cmd
}
