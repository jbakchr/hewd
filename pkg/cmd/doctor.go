package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/jbakchr/hewd/internal/config"
	idoctor "github.com/jbakchr/hewd/internal/doctor"
	iscan "github.com/jbakchr/hewd/internal/scan"
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

			cfg, err := config.Load(cwd)
			if err != nil {
				return err
			}

			summary, err := iscan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			results := idoctor.RunAll(summary, cfg)

			jsonOut, _ := cmd.Flags().GetBool("json")
			yamlOut, _ := cmd.Flags().GetBool("yaml")
			pretty, _ := cmd.Flags().GetBool("pretty")

			if jsonOut && yamlOut {
				return fmt.Errorf("cannot use --json and --yaml together")
			}

			if jsonOut {
				if pretty {
					data, _ := json.MarshalIndent(results, "", "  ")
					fmt.Println(string(data))
				} else {
					data, _ := json.Marshal(results)
					fmt.Println(string(data))
				}
				return nil
			}

			if yamlOut {
				data, _ := yaml.Marshal(results)
				fmt.Println(string(data))
				return nil
			}

			printDoctorResults(results)
			return nil
		},
	}

	cmd.Flags().Bool("json", false, "Output diagnostics in JSON format")
	cmd.Flags().Bool("yaml", false, "Output diagnostics in YAML format")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON output")

	return cmd
}
