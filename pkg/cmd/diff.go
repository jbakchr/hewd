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

			// Validate schema version
			if oldReport.SchemaVersion != newReport.SchemaVersion {
				return fmt.Errorf(
					"schema version mismatch: old=%d new=%d",
					oldReport.SchemaVersion,
					newReport.SchemaVersion,
				)
			}

			// ------------------------------------------------------------
			// PHASE 4: Compute diff + sorting + grouping
			// ------------------------------------------------------------
			result := diff.ComputeDiff(&oldReport, &newReport)

			groupedNew := diff.GroupIssues(result.NewIssues)
			groupedResolved := diff.GroupIssues(result.ResolvedIssues)

			// ------------------------------------------------------------
			// Output
			// ------------------------------------------------------------
			fmt.Println("\n===== OVERALL SCORE =====")
			fmt.Println("----------------------------------------")
			fmt.Printf("Old report score: %d\n", oldReport.Score)
			fmt.Printf("New report score: %d\n", newReport.Score)
			fmt.Printf("Score change: %+d\n", result.ScoreDelta)

			fmt.Println("\n===== CATEGORY SCORES =====")
			fmt.Println("----------------------------------------")

			fmt.Printf("  %-14s %3d → %3d   (%+d)\n",
				"Documentation:",
				oldReport.CategoryScores.Documentation,
				newReport.CategoryScores.Documentation,
				result.CategoryDeltas["documentation"],
			)

			fmt.Printf("  %-14s %3d → %3d   (%+d)\n",
				"Config:",
				oldReport.CategoryScores.Config,
				newReport.CategoryScores.Config,
				result.CategoryDeltas["config"],
			)

			fmt.Printf("  %-14s %3d → %3d   (%+d)\n",
				"Structure:",
				oldReport.CategoryScores.Structure,
				newReport.CategoryScores.Structure,
				result.CategoryDeltas["structure"],
			)

			// ------------------------------------------------------------
			// NEW ISSUES (sorted + grouped)
			// ------------------------------------------------------------
			fmt.Println("\n===== NEW ISSUES =====")
			fmt.Println("----------------------------------------")

			if len(result.NewIssues) == 0 {
				fmt.Println("(none)")
			} else {
				for category, issues := range groupedNew {
					fmt.Printf("\n# %s\n", category)
					for _, issue := range issues {
						fmt.Printf("  - %s (%s)\n", issue.ID, issue.Level)
					}
				}
			}

			// ------------------------------------------------------------
			// RESOLVED ISSUES (sorted + grouped)
			// ------------------------------------------------------------
			fmt.Println("\n===== RESOLVED ISSUES =====")
			fmt.Println("----------------------------------------")

			if len(result.ResolvedIssues) == 0 {
				fmt.Println("(none)")
			} else {
				for category, issues := range groupedResolved {
					fmt.Printf("\n# %s\n", category)
					for _, issue := range issues {
						fmt.Printf("  - %s (%s)\n", issue.ID, issue.Level)
					}
				}
			}

			return nil
		},
	}

	return cmd
}
