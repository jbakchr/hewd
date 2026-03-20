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

	return cmd
}
