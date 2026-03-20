package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jbakchr/hewd/internal/api"
	"github.com/jbakchr/hewd/internal/diff"
)

func newDiffCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "diff <old.json> <new.json>",
		Short: "Compare two hewd JSON reports and show score/regression differences",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			jsonFlag, _ := cmd.Flags().GetBool("json")
			yamlFlag, _ := cmd.Flags().GetBool("yaml")
			mdFlag, _ := cmd.Flags().GetBool("md")

			failScoreDrop, _ := cmd.Flags().GetInt("fail-on-score-drop")
			failNewErrors, _ := cmd.Flags().GetBool("fail-on-new-errors")
			failAny, _ := cmd.Flags().GetBool("fail-on-any-regression")

			oldPath := args[0]
			newPath := args[1]

			// Load old.json
			oldData, err := os.ReadFile(oldPath)
			if err != nil {
				return fmt.Errorf("failed to read old report: %w", err)
			}

			var oldReport api.MachineOutput
			if err := json.Unmarshal(oldData, &oldReport); err != nil {
				return fmt.Errorf("failed to parse old report JSON: %w", err)
			}

			// Load new.json
			newData, err := os.ReadFile(newPath)
			if err != nil {
				return fmt.Errorf("failed to read new report: %w", err)
			}

			var newReport api.MachineOutput
			if err := json.Unmarshal(newData, &newReport); err != nil {
				return fmt.Errorf("failed to parse new report JSON: %w", err)
			}

			if oldReport.SchemaVersion != newReport.SchemaVersion {
				return fmt.Errorf("schema version mismatch: old=%d new=%d",
					oldReport.SchemaVersion, newReport.SchemaVersion)
			}

			// Compute full diff
			result := diff.ComputeDiff(&oldReport, &newReport)

			gate := diff.EvaluateRegressionGates(result, failScoreDrop, failNewErrors, failAny)

			if gate.Failed {
				// Optional: print reasons unless in JSON/YAML mode
				if !jsonFlag && !yamlFlag {
					fmt.Println("\n❌ Regression detected:")
					for _, r := range gate.Reasons {
						fmt.Printf(" - %s\n", r)
					}
				}

				// Exit code 1 means "regression failure"
				return fmt.Errorf("regression gating failed")
			}

			// Make machine-readable structure
			out := diff.MakeDiffOutput(result, &oldReport, &newReport)

			// ------------------------------------------------------------
			// MACHINE READABLE OUTPUT (JSON / YAML)
			// ------------------------------------------------------------
			if jsonFlag {
				return diff.WriteJSON(out)
			}
			if yamlFlag {
				return diff.WriteYAML(out)
			}
			if mdFlag {
				md := diff.WriteMarkdown(result, oldReport, newReport)
				fmt.Println(md)
				return nil
			}

			// ------------------------------------------------------------
			// DEFAULT PRETTY TERMINAL OUTPUT
			// ------------------------------------------------------------
			return writePrettyDiff(result, oldReport, newReport)
		},
	}

	// Add flags
	flags := cmd.Flags()
	flags.Bool("json", false, "Output diff result in JSON format")
	flags.Bool("yaml", false, "Output diff result in YAML format")
	flags.Bool("md", false, "Output diff result in Markdown format")

	flags.Int("fail-on-score-drop", 0,
		"Fail if score drops by N or more points")
	flags.Bool("fail-on-new-errors", false,
		"Fail if any new error-level issues appear")
	flags.Bool("fail-on-any-regression", false,
		"Fail on any regression (score drop, new issues)")

	return cmd
}
