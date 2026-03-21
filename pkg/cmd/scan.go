package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

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

	cmd.GroupID = "analysis"

	cmd.Flags().Bool("json", false, "Output results in JSON format. Use --pretty for indented JSON.")
	cmd.Flags().Bool("yaml", false, "Output results in YAML format.")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON output for readability.")

	return cmd
}
