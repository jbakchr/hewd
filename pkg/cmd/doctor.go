package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/scan"
	"github.com/jbakchr/hewd/internal/score"
)

type DoctorOutput struct {
	Score    int                  `json:"score" yaml:"score"`
	Category score.CategoryScores `json:"category_scores" yaml:"category_scores"`
	Results  []score.ScoredRule   `json:"results" yaml:"results"`
}

func newDoctorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Run diagnostic checks on the project",
		Long: `Runs validation rules against the scanned project structure,
supports categories (--only/--except), scoring (--score), structured output,
and CI-friendly exit codes (--fail-on).`,
		RunE: func(cmd *cobra.Command, args []string) error {

			// Load cwd
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			// Load config (optional)
			cfg, _ := config.Load(cwd)

			// Scan directory
			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			// Flags
			onlyCats, _ := cmd.Flags().GetStringSlice("only")
			exceptCats, _ := cmd.Flags().GetStringSlice("except")
			jsonOut, _ := cmd.Flags().GetBool("json")
			yamlOut, _ := cmd.Flags().GetBool("yaml")
			pretty, _ := cmd.Flags().GetBool("pretty")
			failOnStr, _ := cmd.Flags().GetString("fail-on")
			showScore, _ := cmd.Flags().GetBool("score")
			showCategoryScore, _ := cmd.Flags().GetBool("category-score")

			// Prevent conflict
			if jsonOut && yamlOut {
				return fmt.Errorf("cannot combine --json and --yaml")
			}

			// Run rules
			results := rules.RunAll(summary, cfg, onlyCats, exceptCats)

			// Build score wrapper
			scored := score.ScoreResult{
				Score:   score.Score(results, cfg),
				Results: wrapResultsWithCategory(results),
			}
			// Prepare detailed output structure
			categoryScores := score.ScoreByCategory(scored.Results, cfg)

			out := DoctorOutput{
				Score:    scored.Score,
				Category: categoryScores,
				Results:  scored.Results,
			}

			// ----- JSON -----
			if jsonOut {
				var data []byte
				if pretty {
					data, _ = json.MarshalIndent(out, "", "  ")
				} else {
					data, _ = json.Marshal(out)
				}
				fmt.Println(string(data))
				return evaluateDoctorExitCode(results, failOnStr)
			}

			// ----- YAML -----
			if yamlOut {
				data, _ := yaml.Marshal(out)
				fmt.Println(string(data))
				return evaluateDoctorExitCode(results, failOnStr)
			}

			// ----- PRETTY OUTPUT -----

			if showScore {
				fmt.Printf("Overall Score: %d/100\n", scored.Score)
			}

			if showCategoryScore {
				fmt.Printf("Documentation Score: %d/100\n", categoryScores.Documentation)
				fmt.Printf("Config Score:        %d/100\n", categoryScores.Config)
				fmt.Printf("Structure Score:     %d/100\n", categoryScores.Structure)
				fmt.Printf("Overall Score:       %d/100\n\n", categoryScores.Overall)
			}

			printDoctorPretty(results)

			return evaluateDoctorExitCode(results, failOnStr)
		},
	}

	// Flags
	cmd.Flags().StringSlice("only", []string{}, "Only run rules from these categories (comma-separated)")
	cmd.Flags().StringSlice("except", []string{}, "Exclude rules from these categories (comma-separated)")
	cmd.Flags().Bool("json", false, "Output results in JSON format")
	cmd.Flags().Bool("yaml", false, "Output results in YAML format")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON output")
	cmd.Flags().String("fail-on", "error", "Fail on this severity or higher (info|warn|error)")
	cmd.Flags().Bool("score", false, "Show project maturity score")
	cmd.Flags().Bool("category-score", false, "Show per-category maturity scores")

	return cmd
}

// Pretty-print grouped by category
func printDoctorPretty(results []rules.Result) {
	if len(results) == 0 {
		fmt.Println("No issues found. Project looks healthy!")
		return
	}

	fmt.Println("Doctor Results by Category:")
	grouped := map[string][]rules.Result{}

	for _, r := range results {
		cat := rules.CategoryForRule(r.ID)
		grouped[cat] = append(grouped[cat], r)
	}

	for cat, list := range grouped {
		fmt.Printf("\n[%s]\n", strings.ToUpper(cat))
		for _, r := range list {
			fmt.Printf("  [%s] %s: %s\n", r.Level, r.ID, r.Message)
			if r.File != "" {
				fmt.Printf("    File: %s\n", r.File)
			}
		}
	}
}

// Convert rules.Result → ScoreResult ScoredRule form
func wrapResultsWithCategory(results []rules.Result) []score.ScoredRule {
	out := []score.ScoredRule{}
	for _, r := range results {
		out = append(out, score.ScoredRule{
			Result:   r,
			Category: rules.CategoryForRule(r.ID),
		})
	}
	return out
}

// Exit code logic (fail-on)
func evaluateDoctorExitCode(results []rules.Result, failOn string) error {

	highest := rules.Info
	for _, r := range results {
		if rules.SeverityRank(r.Level) > rules.SeverityRank(highest) {
			highest = r.Level
		}
	}

	var target rules.Level
	switch strings.ToLower(failOn) {
	case "info":
		target = rules.Info
	case "warn":
		target = rules.Warn
	case "error":
		target = rules.Error
	default:
		return fmt.Errorf("invalid value for --fail-on: %s", failOn)
	}

	if rules.SeverityRank(highest) >= rules.SeverityRank(target) {
		switch highest {
		case rules.Error:
			os.Exit(2)
		case rules.Warn:
			os.Exit(1)
		case rules.Info:
			os.Exit(1)
		}
	}

	return nil
}
