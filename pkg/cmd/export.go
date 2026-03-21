package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/jbakchr/hewd/internal/api"
	"github.com/jbakchr/hewd/internal/cliutils"
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
		jsonOut    bool
		yamlOut    bool
		mdOut      bool // markdown export currently not implemented
		prettyJSON bool
	)

	cmd := &cobra.Command{
		Use:     helptext.ExportUse,
		Short:   helptext.ExportShort,
		Long:    helptext.ExportLong,
		Example: helptext.ExportExample,

		RunE: func(cmd *cobra.Command, args []string) error {

			// Ensure output or stdout mode
			if output == "" && !jsonOut && !yamlOut && !mdOut {
				return cliutils.ErrHint(
					"--output is required unless using --json or --yaml to write to stdout",
					"add --output <file> or specify a machine-readable format flag",
				)
			}

			// Validate output flags
			if err := cliutils.ValidateOutputFormatFlags(jsonOut, yamlOut, mdOut, prettyJSON, "export"); err != nil {
				return err
			}

			// Working directory
			cwd, err := os.Getwd()
			if err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to determine working directory: %v", err),
					"ensure the current directory is accessible",
				)
			}

			cfg, err := config.Load(cwd)
			if err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to load configuration: %v", err),
					"ensure .hewd/config.yaml is valid yaml",
				)
			}

			// Scan directory
			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to scan directory: %v", err),
					"ensure this is a valid repository for hewd analysis",
				)
			}

			// Run rules (export always evaluates all categories)
			results := rules.RunAll(summary, cfg, nil, nil)
			scored := score.NewScoredRules(results)

			categoryScores := score.ScoreByCategory(scored, cfg)
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

			// Construct machine-readable output object
			machine := api.MachineOutput{
				SchemaVersion:  1,
				HewdVersion:    version.Version,
				GeneratedAt:    time.Now(),
				Score:          overallScore,
				CategoryScores: categoryScores,
				Results:        scored,
				Fixable:        fixables,
			}

			// -----------------------------------------------------------------
			// JSON output (stdout or file)
			// -----------------------------------------------------------------
			if jsonOut {
				var data []byte
				if prettyJSON {
					data, err = json.MarshalIndent(machine, "", "  ")
				} else {
					data, err = json.Marshal(machine)
				}

				if err != nil {
					return cliutils.ErrHint(
						fmt.Sprintf("failed to marshal json: %v", err),
						"ensure the output is valid json",
					)
				}

				// stdout mode
				if output == "" {
					fmt.Println(string(data))
					return nil
				}

				return writeExportFile(output, data)
			}

			// -----------------------------------------------------------------
			// YAML output (stdout or file)
			// -----------------------------------------------------------------
			if yamlOut {
				data, err := yaml.Marshal(machine)
				if err != nil {
					return cliutils.ErrHint(
						fmt.Sprintf("failed to marshal yaml: %v", err),
						"ensure the output is valid yaml",
					)
				}

				if output == "" {
					fmt.Println(string(data))
					return nil
				}

				return writeExportFile(output, data)
			}

			// -----------------------------------------------------------------
			// Markdown output (not implemented yet)
			// -----------------------------------------------------------------
			if mdOut {
				return cliutils.ErrHint(
					"markdown export is not implemented",
					"use --json or --yaml for machine-readable output",
				)
			}

			// -----------------------------------------------------------------
			// Default: pretty JSON to file
			// -----------------------------------------------------------------
			data, err := json.MarshalIndent(machine, "", "  ")
			if err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to marshal json: %v", err),
					"ensure the output is valid json",
				)
			}

			return writeExportFile(output, data)
		},
	}

	cmd.GroupID = "reporting"

	// Flags
	cmd.Flags().StringVar(&output, "output", "", "Write output to the specified file path (required unless using stdout).")
	cmd.Flags().BoolVar(&yamlOut, "yaml", false, "Export report in YAML format.")
	cmd.Flags().BoolVar(&jsonOut, "json", false, "Export report in JSON format. Use --pretty for indented JSON.")
	cmd.Flags().BoolVar(&mdOut, "md", false, "Export report in Markdown format.")
	cmd.Flags().BoolVar(&prettyJSON, "pretty", false, "Pretty-print JSON output for readability.")

	return cmd
}

// -----------------------------------------------------------------------------
// Helper for writing output files safely.
// -----------------------------------------------------------------------------
func writeExportFile(path string, data []byte) error {
	dir := filepath.Dir(path)
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return cliutils.ErrHint(
				fmt.Sprintf("failed to create directory %s: %v", dir, err),
				"ensure parent directories are writable",
			)
		}
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return cliutils.ErrHint(
			fmt.Sprintf("failed to write file %s: %v", path, err),
			"ensure directory permissions allow writing",
		)
	}

	fmt.Printf("Export complete → %s\n", path)
	return nil
}
