package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/scan"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newDoctorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Run diagnostic checks on the project",
		RunE: func(cmd *cobra.Command, args []string) error {

			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			results := rules.RunAll(summary)

			jsonOut, _ := cmd.Flags().GetBool("json")
			yamlOut, _ := cmd.Flags().GetBool("yaml")
			pretty, _ := cmd.Flags().GetBool("pretty")

			if jsonOut && yamlOut {
				return fmt.Errorf("cannot use --json and --yaml together")
			}

			if jsonOut {
				var data []byte
				if pretty {
					data, _ = json.MarshalIndent(results, "", "  ")
				} else {
					data, _ = json.Marshal(results)
				}
				fmt.Println(string(data))
				return nil
			}

			if yamlOut {
				data, _ := yaml.Marshal(results)
				fmt.Println(string(data))
				return nil
			}

			if len(results) == 0 {
				fmt.Println("No issues found. Project looks healthy!")
				return nil
			}

			fmt.Println("Doctor Results:")
			for _, r := range results {
				fmt.Printf("[%s] %s: %s\n", r.Level, r.ID, r.Message)
				if r.File != "" {
					fmt.Printf("  File: %s\n", r.File)
				}
			}

			return nil
		},
	}

	cmd.Flags().Bool("json", false, "Output results in JSON format")
	cmd.Flags().Bool("yaml", false, "Output results in YAML format")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON output")

	return cmd
}
