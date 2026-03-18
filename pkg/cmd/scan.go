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
		Long:  `Scan the current directory, detecting languages, documentation, and config files.`,
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

	cmd.Flags().Bool("json", false, "Output scan result in JSON format")
	cmd.Flags().Bool("yaml", false, "Output scan result in YAML format")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON output")

	return cmd
}
