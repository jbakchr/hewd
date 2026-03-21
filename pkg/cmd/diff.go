package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jbakchr/hewd/internal/api"
	"github.com/jbakchr/hewd/internal/diff"
	"github.com/jbakchr/hewd/internal/helptext"
)

func newDiffCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     helptext.DiffUse,
		Short:   helptext.DiffShort,
		Long:    helptext.DiffLong,
		Example: helptext.DiffExample,
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			// --------------------------------------
			// Flags
			// --------------------------------------
			jsonFlag, _ := cmd.Flags().GetBool("json")
			yamlFlag, _ := cmd.Flags().GetBool("yaml")
			mdFlag, _ := cmd.Flags().GetBool("md")

			failScoreDrop, _ := cmd.Flags().GetInt("fail-on-score-drop")
			failNewErrors, _ := cmd.Flags().GetBool("fail-on-new-errors")
			failAny, _ := cmd.Flags().GetBool("fail-on-any-regression")

			// Prevent invalid format combos
			if (jsonFlag && yamlFlag) ||
				(jsonFlag && mdFlag) ||
				(yamlFlag && mdFlag) {
				return fmt.Errorf("cannot combine --json, --yaml, or --md")
			}

			oldPath := args[0]
			newPath := args[1]

			// --------------------------------------
			// Load old report
			// --------------------------------------
			oldData, err := os.ReadFile(oldPath)
			if err != nil {
				return fmt.Errorf("failed to read old report: %w", err)
			}

			var oldReport api.MachineOutput
			if err := json.Unmarshal(oldData, &oldReport); err != nil {
				return fmt.Errorf("failed to parse old report JSON: %w", err)
			}

			// --------------------------------------
			// Load new report
			// --------------------------------------
			newData, err := os.ReadFile(newPath)
			if err != nil {
				return fmt.Errorf("failed to read new report: %w", err)
			}

			var newReport api.MachineOutput
			if err := json.Unmarshal(newData, &newReport); err != nil {
				return fmt.Errorf("failed to parse new report JSON: %w", err)
			}

			// Schema version check
			if oldReport.SchemaVersion != newReport.SchemaVersion {
				return fmt.Errorf(
					"schema version mismatch: old=%d new=%d",
					oldReport.SchemaVersion, newReport.SchemaVersion,
				)
			}

			// --------------------------------------
			// Compute diff
			// --------------------------------------
			result := diff.ComputeDiff(&oldReport, &newReport)

			// Regression gating
			gate := diff.EvaluateRegressionGates(result, failScoreDrop, failNewErrors, failAny)

			if gate.Failed {
				if !jsonFlag && !yamlFlag {
					fmt.Println("\n❌ Regression detected:")
					for _, r := range gate.Reasons {
						fmt.Printf(" - %s\n", r)
					}
				}
				return fmt.Errorf("regression gating failed")
			}

			// Machine-readable diff structure
			out := diff.MakeDiffOutput(result, &oldReport, &newReport)

			// --------------------------------------
			// JSON / YAML / Markdown outputs
			// --------------------------------------
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

			// --------------------------------------
			// DEFAULT: Pretty terminal diff
			// --------------------------------------
			return writePrettyDiff(result, oldReport, newReport)
		},
	}

	// ----- Command Group -----
	cmd.GroupID = "analysis"

	// --------------------------------------
	// Flags
	// --------------------------------------
	cmd.Flags().Bool("json", false, "Output diff results in JSON format. Use --pretty for indented JSON.")
	cmd.Flags().Bool("yaml", false, "Output diff results in YAML format.")
	cmd.Flags().Bool("md", false, "Output diff results in Markdown format.")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON output for readability.")

	cmd.Flags().Int("fail-on-score-drop", 0, "Fail if score drops by N or more points.")
	cmd.Flags().Bool("fail-on-new-errors", false, "Fail if any new error-level issues appear.")
	cmd.Flags().Bool("fail-on-any-regression", false, "Fail on any regression (score drop or new issues).")

	return cmd
}
