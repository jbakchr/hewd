package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jbakchr/hewd/internal/api"
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

			// Validate files exist (logic only, no diffing yet)
			if _, err := os.Stat(oldPath); os.IsNotExist(err) {
				return fmt.Errorf("old report not found: %s", oldPath)
			}
			if _, err := os.Stat(newPath); os.IsNotExist(err) {
				return fmt.Errorf("new report not found: %s", newPath)
			}

			fmt.Println("\n===== OVERALL SCORE =====")
			fmt.Printf("Old report score: %d\n", oldReport.Score)
			fmt.Printf("New report score: %d\n", newReport.Score)

			// Score delta
			delta := newReport.Score - oldReport.Score
			fmt.Printf("Score change: %+d\n", delta)

			fmt.Println("\n===== CATEGORY SCORES =====")
			// Category score deltas
			docDelta := newReport.CategoryScores.Documentation - oldReport.CategoryScores.Documentation
			cfgDelta := newReport.CategoryScores.Config - oldReport.CategoryScores.Config
			structDelta := newReport.CategoryScores.Structure - oldReport.CategoryScores.Structure

			fmt.Println()
			fmt.Println("Category changes:")
			fmt.Printf("  Documentation: %d → %d (%+d)\n",
				oldReport.CategoryScores.Documentation,
				newReport.CategoryScores.Documentation,
				docDelta,
			)
			fmt.Printf("  Config:        %d → %d (%+d)\n",
				oldReport.CategoryScores.Config,
				newReport.CategoryScores.Config,
				cfgDelta,
			)
			fmt.Printf("  Structure:     %d → %d (%+d)\n",
				oldReport.CategoryScores.Structure,
				newReport.CategoryScores.Structure,
				structDelta,
			)

			// --------------------------------------------
			// New issues (present in newReport but not oldReport)
			// --------------------------------------------
			// Build set of old issue IDs
			oldIDs := make(map[string]bool)
			for _, r := range oldReport.Results {
				oldIDs[r.ID] = true
			}

			fmt.Println("\n===== NEW ISSUES =====")
			fmt.Println("New issues:")
			foundNew := false
			for _, r := range newReport.Results {
				if !oldIDs[r.ID] {
					fmt.Printf("  - %s (%s)\n", r.ID, r.Level)
					foundNew = true
				}
			}
			if !foundNew {
				fmt.Println("  (none)")
			}

			// --------------------------------------------
			// Resolved issues (present in oldReport but not in newReport)
			// --------------------------------------------
			newIDs := make(map[string]bool)
			for _, r := range newReport.Results {
				newIDs[r.ID] = true
			}

			fmt.Println("\n===== RESOLVED ISSUES =====")
			fmt.Println("Resolved issues:")
			foundResolved := false
			for _, r := range oldReport.Results {
				if !newIDs[r.ID] {
					fmt.Printf("  - %s (%s)\n", r.ID, r.Level)
					foundResolved = true
				}
			}
			if !foundResolved {
				fmt.Println("  (none)")
			}

			return nil
		},
	}

	return cmd
}
