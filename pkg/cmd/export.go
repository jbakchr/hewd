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
		Use:   "export",
		Short: "Export a complete machine-readable project health report.",
		Long: `hewd export generates a complete, machine-readable report describing the
documentation, configuration, and structural health of the current repository.

The exported report follows hewd's stable MachineOutput schema, which includes:

  • Overall project health score
  • Category scores (documentation, config, structure)
  • Detailed rule results with severity levels
  • Fixable items
  • Version metadata
  • Timestamps

Exported reports are ideal for CI pipelines, dashboards, trend tracking, and
use as input for 'hewd diff'. The command uses the same diagnostic engine as
'hewd doctor' but omits human-focused formatting in favor of stable,
automation-friendly output structures.`,
		Example: `
  # Export project health to JSON
  hewd export --output hewd.json

  # Export in YAML format
  hewd export --yaml --output hewd.yaml

  # Pretty-print JSON to stdout
  hewd export --json --pretty

  # Use export to generate reports for diff comparison
  hewd export --output old.json
  # ... make changes ...
  hewd export --output new.json
  hewd diff old.json new.json

  # Pipe export data to another command
  hewd export --json | jq '.score'
`,
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

	// ----- Flags -----
	cmd.Flags().StringVar(&output, "output", "", "Path to write machine-readable output (JSON by default)")
	cmd.Flags().BoolVar(&yamlOut, "yaml", false, "Export report as YAML (writes to stdout if --output is not set)")
	cmd.Flags().BoolVar(&jsonOut, "json", false, "Export report as JSON (writes to stdout if --output is not set)")
	cmd.Flags().BoolVar(&prettyJSON, "pretty", false, "Pretty-print JSON output")

	return cmd
}

func writeExportFile(path string, data []byte) error {
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}
	fmt.Printf("Export complete → %s\n", path)
	return nil
}
