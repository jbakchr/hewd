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
		yamlOut    bool
		jsonOut    bool
		mdOut      bool // currently unsupported, but included for consistency/future
		prettyJSON bool
	)

	cmd := &cobra.Command{
		Use:     helptext.ExportUse,
		Short:   helptext.ExportShort,
		Long:    helptext.ExportLong,
		Example: helptext.ExportExample,

		RunE: func(cmd *cobra.Command, args []string) error {

			// If no output file is provided AND stdout formats are not used, error.
			if output == "" && !jsonOut && !yamlOut && !mdOut {
				return fmt.Errorf("--output is required unless using --json or --yaml directly to stdout")
			}

			// Validate JSON/YAML/MD/pretty flags using shared helper
			if err := cliutils.ValidateOutputFormatFlags(jsonOut, yamlOut, mdOut, prettyJSON, "hewd export"); err != nil {
				return err
			}

			// Determine working directory
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not determine working directory: %w", err)
			}

			cfg, _ := config.Load(cwd)

			// Scan the repository
			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			// Run all rules (export always includes full ruleset)
			results := rules.RunAll(summary, cfg, nil, nil)
			scoredRules := score.NewScoredRules(results)

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

			// Construct machine‑readable output object
			machine := api.MachineOutput{
				SchemaVersion:  1,
				HewdVersion:    version.Version,
				GeneratedAt:    time.Now(),
				Score:          overallScore,
				CategoryScores: categoryScores,
				Results:        scoredRules,
				Fixable:        fixables,
			}

			// ------------------------------
			// YAML output (stdout or file)
			// ------------------------------
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

			// ------------------------------
			// JSON output (stdout or file)
			// ------------------------------
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

				// stdout
				if output == "" {
					fmt.Println(string(data))
					return nil
				}

				return writeExportFile(output, data)
			}

			// ------------------------------
			// Markdown output (future)
			// ------------------------------
			if mdOut {
				return fmt.Errorf("Markdown output is not yet implemented for `hewd export`")
			}

			// ------------------------------
			// Default: JSON to file
			// ------------------------------
			data, err := json.MarshalIndent(machine, "", "  ")
			if err != nil {
				return err
			}

			return writeExportFile(output, data)
		},
	}

	cmd.GroupID = "reporting"

	// Output flags (consistent ordering)
	cmd.Flags().BoolVar(&yamlOut, "yaml", false, "Export report in YAML format.")
	cmd.Flags().BoolVar(&jsonOut, "json", false, "Export report in JSON format. Use --pretty for indented JSON.")
	cmd.Flags().BoolVar(&mdOut, "md", false, "Export report in Markdown format.") // Optional future support
	cmd.Flags().BoolVar(&prettyJSON, "pretty", false, "Pretty-print JSON output for readability.")

	// Output path
	cmd.Flags().StringVar(&output, "output", "", "Write output to the specified file path (required unless using stdout).")

	return cmd
}

func writeExportFile(path string, data []byte) error {
	// Ensure parent directories exist
	if dir := filepath.Dir(path); dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}

	fmt.Printf("Export complete → %s\n", path)
	return nil
}
