package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

func newDoctorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     helptext.DoctorUse,
		Short:   helptext.DoctorShort,
		Long:    helptext.DoctorLong,
		Example: helptext.DoctorExample,

		RunE: func(cmd *cobra.Command, args []string) error {

			// -----------------------------------------------------------------
			// Working directory
			// -----------------------------------------------------------------
			cwd, err := os.Getwd()
			if err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to determine working directory: %v", err),
					"ensure the current directory is accessible",
				)
			}

			// -----------------------------------------------------------------
			// Load config
			// -----------------------------------------------------------------
			cfg, err := config.Load(cwd)
			if err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to load configuration: %v", err),
					"ensure .hewd/config.yaml is valid yaml and readable",
				)
			}

			// -----------------------------------------------------------------
			// Scan project
			// -----------------------------------------------------------------
			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return cliutils.ErrHint(
					fmt.Sprintf("failed to scan directory: %v", err),
					"ensure you are running hewd inside a valid repository",
				)
			}

			// -----------------------------------------------------------------
			// Parse flags
			// -----------------------------------------------------------------
			jsonOut, _ := cmd.Flags().GetBool("json")
			yamlOut, _ := cmd.Flags().GetBool("yaml")
			mdOut, _ := cmd.Flags().GetBool("md")
			pretty, _ := cmd.Flags().GetBool("pretty")

			onlyCats, _ := cmd.Flags().GetStringSlice("only")
			exceptCats, _ := cmd.Flags().GetStringSlice("except")

			showScore, _ := cmd.Flags().GetBool("score")
			showCategoryScore, _ := cmd.Flags().GetBool("category-score")

			failOnStr, _ := cmd.Flags().GetString("fail-on")

			// -----------------------------------------------------------------
			// Validate output flags
			// -----------------------------------------------------------------
			if err := cliutils.ValidateOutputFormatFlags(jsonOut, yamlOut, mdOut, pretty, "doctor"); err != nil {
				return err
			}

			if (showScore || showCategoryScore) && (jsonOut || yamlOut || mdOut) {
				return cliutils.ErrHint(
					"--score or --category-score cannot be combined with machine-readable output flags",
					"remove --score or --category-score or omit --json/--yaml/--md",
				)
			}

			// -----------------------------------------------------------------
			// Run rules with category filtering
			// -----------------------------------------------------------------
			results := rules.RunAll(summary, cfg, onlyCats, exceptCats)

			// -----------------------------------------------------------------
			// Scored rules + category scores
			// -----------------------------------------------------------------
			scoredRules := score.NewScoredRules(results)

			categoryScores := score.ScoreByCategory(scoredRules, cfg)
			overallScore := score.Score(results, cfg)

			// -----------------------------------------------------------------
			// Detect fixables
			// -----------------------------------------------------------------
			rawFixes := fix.DetectFixes(results, cwd)
			var fixables []api.FixableItem

			for _, f := range rawFixes {
				fixables = append(fixables, api.FixableItem{
					RuleID:   f.RuleID,
					Message:  f.Message,
					FilePath: f.FilePath,
				})
			}

			// -----------------------------------------------------------------
			// Construct machine-readable output object
			// -----------------------------------------------------------------
			machine := api.MachineOutput{
				SchemaVersion:  1,
				HewdVersion:    version.Version,
				GeneratedAt:    time.Now(),
				Score:          overallScore,
				CategoryScores: categoryScores,
				Results:        scoredRules,
				Fixable:        fixables,
			}

			// -----------------------------------------------------------------
			// Markdown output
			// -----------------------------------------------------------------
			if mdOut {
				md := renderMarkdown(machine)
				fmt.Println(md)
				return evaluateDoctorExitCode(results, failOnStr)
			}

			// -----------------------------------------------------------------
			// JSON output
			// -----------------------------------------------------------------
			if jsonOut {
				var data []byte
				if pretty {
					data, err = json.MarshalIndent(machine, "", "  ")
				} else {
					data, err = json.Marshal(machine)
				}

				if err != nil {
					return cliutils.ErrHint(
						fmt.Sprintf("failed to marshal json: %v", err),
						"ensure json output is well-formed",
					)
				}

				fmt.Println(string(data))
				return evaluateDoctorExitCode(results, failOnStr)
			}

			// -----------------------------------------------------------------
			// YAML output
			// -----------------------------------------------------------------
			if yamlOut {
				data, err := yaml.Marshal(machine)
				if err != nil {
					return cliutils.ErrHint(
						fmt.Sprintf("failed to marshal yaml: %v", err),
						"ensure yaml output is well-formed",
					)
				}

				fmt.Println(string(data))
				return evaluateDoctorExitCode(results, failOnStr)
			}

			// -----------------------------------------------------------------
			// Pretty Output
			// -----------------------------------------------------------------
			if showScore {
				fmt.Printf("Overall Score: %d/100\n", overallScore)
			}

			if showCategoryScore {
				fmt.Printf("Documentation Score: %d\n", categoryScores.Documentation)
				fmt.Printf("Config Score:        %d\n", categoryScores.Config)
				fmt.Printf("Structure Score:     %d\n", categoryScores.Structure)
				fmt.Printf("Overall:             %d\n\n", categoryScores.Overall)
			}

			printDoctorPretty(unwrapResults(scoredRules))

			// -----------------------------------------------------------------
			// CI gating
			// -----------------------------------------------------------------
			return evaluateDoctorExitCode(results, failOnStr)
		},
	}

	// -------------------------------------------------------------------------
	// Flags
	// -------------------------------------------------------------------------
	cmd.GroupID = "analysis"

	cmd.Flags().Bool("json", false, "Output the diagnostic report in JSON format. Use --pretty for indented JSON.")
	cmd.Flags().Bool("yaml", false, "Output the diagnostic report in YAML format.")
	cmd.Flags().Bool("md", false, "Output the diagnostic report in Markdown format.")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON output for readability.")

	cmd.Flags().StringSlice("only", []string{}, "Only run rules from specific categories (comma-separated).")
	cmd.Flags().StringSlice("except", []string{}, "Skip rules from specific categories (comma-separated).")

	cmd.Flags().Bool("score", false, "Print only the overall score (terminal output only).")
	cmd.Flags().Bool("category-score", false, "Print only category scores (terminal output only).")

	cmd.Flags().String("fail-on", "error", "Fail if a rule of this severity or higher occurs (info|warn|error).")

	return cmd
}

// -----------------------------------------------------------------------------
// Helper: unwrapResults
// -----------------------------------------------------------------------------
func unwrapResults(scored []score.ScoredRule) []rules.Result {
	out := make([]rules.Result, 0, len(scored))
	for _, sr := range scored {
		out = append(out, sr.Result)
	}
	return out
}

// -----------------------------------------------------------------------------
// Helper: evaluateDoctorExitCode
// Applies CI gating rules based on --fail-on severity threshold.
// -----------------------------------------------------------------------------
func evaluateDoctorExitCode(results []rules.Result, failOn string) error {

	// Normalize severity level
	level := strings.ToLower(strings.TrimSpace(failOn))
	if level != "info" && level != "warn" && level != "error" {
		return cliutils.ErrHint(
			fmt.Sprintf("invalid value for --fail-on: %s", failOn),
			"valid values are: info, warn, error",
		)
	}

	// Determine threshold
	var minRank int
	switch level {
	case "info":
		minRank = 1
	case "warn":
		minRank = 2
	case "error":
		minRank = 3
	}

	// Scan for rule results meeting or exceeding threshold
	for _, r := range results {
		if rules.SeverityRank(r.Level) >= minRank {
			return cliutils.ErrHint(
				fmt.Sprintf("diagnostics failed due to %s-level issues", level),
				"address reported issues or choose a less strict --fail-on value",
			)
		}
	}

	// No issues exceeded the threshold → success
	return nil
}
