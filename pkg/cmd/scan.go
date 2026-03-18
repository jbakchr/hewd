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
		Short: "Scan the current directory and summarize the project",
		Long: `Scan the current working directory and produce a project summary,
respecting include/exclude rules from .hewd/config.yaml.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			// Determine project root
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not get working directory: %w", err)
			}

			// Execute scan
			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			// Flags
			jsonOut, _ := cmd.Flags().GetBool("json")
			yamlOut, _ := cmd.Flags().GetBool("yaml")
			pretty, _ := cmd.Flags().GetBool("pretty")

			// Prevent incompatible outputs
			if jsonOut && yamlOut {
				return fmt.Errorf("cannot use --json and --yaml together")
			}

			// -----------------------------------------------------------
			// JSON Output
			// -----------------------------------------------------------
			if jsonOut {
				var data []byte
				if pretty {
					data, err = json.MarshalIndent(summary, "", "  ")
				} else {
					data, err = json.Marshal(summary)
				}
				if err != nil {
					return fmt.Errorf("failed to encode JSON: %w", err)
				}
				fmt.Println(string(data))
				return nil
			}

			// -----------------------------------------------------------
			// YAML Output
			// -----------------------------------------------------------
			if yamlOut {
				data, err := yaml.Marshal(summary)
				if err != nil {
					return fmt.Errorf("failed to encode YAML: %w", err)
				}
				fmt.Println(string(data))
				return nil
			}

			// -----------------------------------------------------------
			// Pretty (default) Output
			// -----------------------------------------------------------
			printScanSummary(summary)
			return nil
		},
	}

	cmd.Flags().Bool("json", false, "Output scan result in JSON format")
	cmd.Flags().Bool("yaml", false, "Output scan result in YAML format")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON output (when --json is used)")

	return cmd
}
