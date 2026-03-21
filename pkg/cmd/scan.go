package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/jbakchr/hewd/internal/cliutils"
	"github.com/jbakchr/hewd/internal/helptext"
	"github.com/jbakchr/hewd/internal/scan"
)

func newScanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     helptext.ScanUse,
		Short:   helptext.ScanShort,
		Long:    helptext.ScanLong,
		Example: helptext.ScanExample,

		RunE: func(cmd *cobra.Command, args []string) error {

			cwd, err := os.Getwd()
			if err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to determine working directory: %v", err),
					"ensure the current directory is accessible",
				)
			}

			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to scan directory: %v", err),
					"ensure you are running hewd inside a valid repository",
				)
			}

			// ---- Parse flags ----
			jsonOut, _ := cmd.Flags().GetBool("json")
			yamlOut, _ := cmd.Flags().GetBool("yaml")
			mdOut, _ := cmd.Flags().GetBool("md")
			pretty, _ := cmd.Flags().GetBool("pretty")

			// ---- Validate output flags ----
			if err := cliutils.ValidateOutputFormatFlags(jsonOut, yamlOut, mdOut, pretty, "scan"); err != nil {
				return err
			}

			// ---- JSON Output ----
			if jsonOut {
				var data []byte
				if pretty {
					data, err = json.MarshalIndent(summary, "", "  ")
				} else {
					data, err = json.Marshal(summary)
				}
				if err != nil {
					return cliutils.ErrHint(
						fmt.Sprintf("failed to marshal json: %v", err),
						"ensure the output structure is valid json",
					)
				}
				fmt.Println(string(data))
				return nil
			}

			// ---- YAML Output ----
			if yamlOut {
				data, err := yaml.Marshal(summary)
				if err != nil {
					return cliutils.ErrHint(
						fmt.Sprintf("failed to marshal yaml: %v", err),
						"ensure the output structure is valid yaml",
					)
				}
				fmt.Println(string(data))
				return nil
			}

			// ---- Markdown (not implemented yet) ----
			if mdOut {
				return cliutils.ErrHint(
					"markdown output is not yet implemented for scan",
					"use --json or --yaml for machine-readable output",
				)
			}

			// ---- Pretty output (default) ----
			printScanSummary(summary)
			return nil
		},
	}

	cmd.GroupID = "analysis"

	cmd.Flags().Bool("json", false, "Output results in JSON format. Use --pretty for indented JSON.")
	cmd.Flags().Bool("yaml", false, "Output results in YAML format.")
	cmd.Flags().Bool("md", false, "Output results in Markdown format.")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON output for readability.")

	return cmd
}
