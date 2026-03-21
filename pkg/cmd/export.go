package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/jbakchr/hewd/internal/api"
	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/fix"
	"github.com/jbakchr/hewd/internal/helptext"
	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/scan"
	"github.com/jbakchr/hewd/internal/score"
	"github.com/jbakchr/hewd/internal/version"
)

func newExportCmd() *cobra.Command {
	var (
		output     string
		yamlOut    bool
		jsonOut    bool
		prettyJSON bool
	)

	cmd := &cobra.Command{
		Use:     helptext.ExportUse,
		Short:   helptext.ExportShort,
		Long:    helptext.ExportLong,
		Example: helptext.ExportExample,
		RunE: func(cmd *cobra.Command, args []string) error {

			// ----- Required flag -----
			if output == "" && !jsonOut && !yamlOut {
				return fmt.Errorf("--output is required unless using --json or --yaml directly to stdout")
			}

			// ----- Format conflicts -----
			if jsonOut && yamlOut {
				return fmt.Errorf("cannot combine --json and --yaml")
			}

			if yamlOut && prettyJSON {
				return fmt.Errorf("cannot combine --yaml and --pretty (pretty applies only to JSON)")
			}

			// ----- Directory -----
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not determine working directory: %w", err)
			}

			cfg, _ := config.Load(cwd)

			// ----- Scan project -----
			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			// ----- Run all rules -----
			results := rules.RunAll(summary, cfg, nil, nil)
			scoredRules := score.NewScoredRules(results)

			categoryScores := score.ScoreByCategory(scoredRules, cfg)
			overallScore := score.Score(results, cfg)

			// ----- Fixable items -----
			rawFixes := fix.DetectFixes(results, cwd)
			var fixables []api.FixableItem
			for _, f := range rawFixes {
				fixables = append(fixables, api.FixableItem{
					RuleID:   f.RuleID,
					Message:  f.Message,
					FilePath: f.FilePath,
				})
			}

			// ----- Build machine output -----
			machine := api.MachineOutput{
				SchemaVersion:  1,
				HewdVersion:    version.Version,
				GeneratedAt:    time.Now(),
				Score:          overallScore,
				CategoryScores: categoryScores,
				Results:        scoredRules,
				Fixable:        fixables,
			}

			// ----- YAML Output -----
			if yamlOut {
				data, err := yaml.Marshal(machine)
				if err != nil {
					return err
				}

				// stdout mode
				if output == "" {
					fmt.Println(string(data))
					return nil
				}

				return writeExportFile(output, data)
			}

			// ----- JSON Output -----
			if jsonOut {
				var data []byte
				if prettyJSON {
					data, err = json.MarshalIndent(machine, "", "  ")
				} else {
					data, err = json.Marshal(machine)
				}
				if err != nil {
					return err
				}

				if output == "" {
					fmt.Println(string(data))
					return nil
				}

				return writeExportFile(output, data)
			}

			// ----- Default: JSON to file -----
			data, err := json.MarshalIndent(machine, "", "  ")
			if err != nil {
				return err
			}

			return writeExportFile(output, data)
		},
	}

	// ----- Command Group -----
	cmd.GroupID = "reporting"

	// ----- Flags -----

	cmd.Flags().StringVar(&output, "output", "", "Write output to the specified file path (required unless using stdout).")

	cmd.Flags().BoolVar(&yamlOut, "yaml", false, "Export report in YAML format.")
	cmd.Flags().BoolVar(&jsonOut, "json", false, "Export report in JSON format. Use --pretty for indented JSON.")
	cmd.Flags().BoolVar(&prettyJSON, "pretty", false, "Pretty-print JSON output for readability.")

	return cmd
}

func writeExportFile(path string, data []byte) error {
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}
	fmt.Printf("Export complete → %s\n", path)
	return nil
}
