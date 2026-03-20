package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/jbakchr/hewd/internal/api"
	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/fix"
	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/scan"
	"github.com/jbakchr/hewd/internal/score"
	"github.com/jbakchr/hewd/internal/version"
)

// newExportCmd provides:
//
//	hewd export --output hewd.json
//
// This produces a stable, versioned machine-readable JSON file containing:
// - overall score
// - category scores
// - full rule results
// - fixable items
// - metadata (version, timestamp)
func newExportCmd() *cobra.Command {
	var output string

	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export a complete machine-readable project health report.",
		Long: `hewd export generates a complete, machine-readable JSON or YAML report
describing the documentation, configuration, and structural health of the 
current repository. The exported output uses hewd's stable MachineOutput schema, 
which includes overall scores, category scores, rule results, fixable items, 
metadata, timestamps, and version information.

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

			if output == "" {
				return fmt.Errorf("--output is required (example: hewd export --output hewd.json)")
			}

			cwd, _ := os.Getwd()
			cfg, _ := config.Load(cwd)

			// Scan project structure
			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			// Run all rules (no category filtering here; export is always full data)
			results := rules.RunAll(summary, cfg, nil, nil)

			// Wrap rule results with category
			scoredRules := score.NewScoredRules(results)

			// Scores
			categoryScores := score.ScoreByCategory(scoredRules, cfg)
			overallScore := score.Score(results, cfg)

			// Fixable items
			rawFixes := fix.DetectFixes(results, cwd)
			var fixables []api.FixableItem
			for _, f := range rawFixes {
				fixables = append(fixables, api.FixableItem{
					RuleID:   f.RuleID,
					Message:  f.Message,
					FilePath: f.FilePath,
				})
			}

			// Build the machine-readable API output
			machine := api.MachineOutput{
				SchemaVersion:  1,
				HewdVersion:    version.Version,
				GeneratedAt:    time.Now(),
				Score:          overallScore,
				CategoryScores: categoryScores,
				Results:        scoredRules,
				Fixable:        fixables,
			}

			// JSON encoding
			data, err := json.MarshalIndent(machine, "", "  ")
			if err != nil {
				return err
			}

			// Write file
			if err := os.WriteFile(output, data, 0644); err != nil {
				return err
			}

			fmt.Printf("Export complete → %s\n", output)
			return nil
		},
	}

	cmd.Flags().StringVar(&output, "output", "", "Path to write machine-readable JSON (required)")
	return cmd
}
