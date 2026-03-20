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

This command is ideal for getting a high-level overview of a project's
documentation and structure before running 'hewd doctor' or exporting a
machine-readable health report.

Scan output can be printed in pretty, JSON, or YAML formats and is safe to run
in local development or CI environments.`,
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

	cmd.Flags().Bool("json", false, "Output full scan results in JSON format (useful for CI and automation)")
	cmd.Flags().Bool("yaml", false, "Output results in YAML format")
	cmd.Flags().Bool("pretty", false, "Show readable, color-friendly output ideal for local development")

	return cmd
}
