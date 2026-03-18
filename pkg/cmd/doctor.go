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
	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/fix"
	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/scan"
	"github.com/jbakchr/hewd/internal/score"
	"github.com/jbakchr/hewd/internal/version"
)

func newDoctorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Run diagnostic checks on the project",
		Long: `Run validation rules against the scanned project structure.
Supports filtering (--only/--except), JSON/YAML/Markdown output, scoring,
and CI-friendly exit codes.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			cfg, _ := config.Load(cwd)

			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			// Flags
			onlyCats, _ := cmd.Flags().GetStringSlice("only")
			exceptCats, _ := cmd.Flags().GetStringSlice("except")
			jsonOut, _ := cmd.Flags().GetBool("json")
			yamlOut, _ := cmd.Flags().GetBool("yaml")
			prettyJSON, _ := cmd.Flags().GetBool("pretty")
			failOnStr, _ := cmd.Flags().GetString("fail-on")
			showScore, _ := cmd.Flags().GetBool("score")
			showCategoryScore, _ := cmd.Flags().GetBool("category-score")
			markdownOut, _ := cmd.Flags().GetBool("md")

			if jsonOut && yamlOut {
				return fmt.Errorf("cannot use --json and --yaml together")
			}

			// Run rules
			results := rules.RunAll(summary, cfg, onlyCats, exceptCats)

			// Wrap rule results as ScoredRule
			scoredRules := score.NewScoredRules(results)

			// Compute scores
			categoryScores := score.ScoreByCategory(scoredRules, cfg)
			overallScore := score.Score(results, cfg)

			// Detect fixable items
			rawFixes := fix.DetectFixes(results, cwd)
			var fixables []api.FixableItem
			for _, f := range rawFixes {
				fixables = append(fixables, api.FixableItem{
					RuleID:   f.RuleID,
					Message:  f.Message,
					FilePath: f.FilePath,
				})
			}

			// Machine-readable output object
			machine := api.MachineOutput{
				SchemaVersion:  1,
				HewdVersion:    version.Version,
				GeneratedAt:    time.Now(),
				Score:          overallScore,
				CategoryScores: categoryScores,
				Results:        scoredRules,
				Fixable:        fixables,
			}

			// ----- Markdown output -----
			if markdownOut {
				md := renderMarkdown(machine)
				fmt.Println(md)
				return evaluateDoctorExitCode(results, failOnStr)
			}

			// ----- JSON output -----
			if jsonOut {
				var data []byte
				if prettyJSON {
					data, _ = json.MarshalIndent(machine, "", "  ")
				} else {
					data, _ = json.Marshal(machine)
				}
				fmt.Println(string(data))
				return evaluateDoctorExitCode(results, failOnStr)
			}

			// ----- YAML output -----
			if yamlOut {
				data, _ := yaml.Marshal(machine)
				fmt.Println(string(data))
				return evaluateDoctorExitCode(results, failOnStr)
			}

			// ----- Pretty Output (terminal) -----
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

			return evaluateDoctorExitCode(results, failOnStr)
		},
	}

	// Flags
	cmd.Flags().StringSlice("only", []string{}, "Only run rules from these categories (comma-separated)")
	cmd.Flags().StringSlice("except", []string{}, "Exclude rules from these categories (comma-separated)")
	cmd.Flags().Bool("json", false, "Output machine-readable JSON")
	cmd.Flags().Bool("yaml", false, "Output machine-readable YAML")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON")
	cmd.Flags().String("fail-on", "error", "Fail on this severity or above (info|warn|error)")
	cmd.Flags().Bool("score", false, "Show overall project score")
	cmd.Flags().Bool("category-score", false, "Show breakdown of category scores")
	cmd.Flags().Bool("md", false, "Output detailed Markdown report")

	return cmd
}

//
// Pretty Output (grouped by category)
//

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

//
// Markdown Output
//

func renderMarkdown(m api.MachineOutput) string {
	var b strings.Builder

	fmt.Fprintf(&b, "# hewd Report\n\n")

	fmt.Fprintf(&b, "Generated by **hewd v%s** on %s\n\n",
		m.HewdVersion, m.GeneratedAt.Format(time.RFC3339))

	// Scores
	fmt.Fprintf(&b, "## Scores\n")
	fmt.Fprintf(&b, "- **Overall:** %d/100\n", m.Score)
	fmt.Fprintf(&b, "- **Documentation:** %d\n", m.CategoryScores.Documentation)
	fmt.Fprintf(&b, "- **Config:** %d\n", m.CategoryScores.Config)
	fmt.Fprintf(&b, "- **Structure:** %d\n\n", m.CategoryScores.Structure)

	fmt.Fprintf(&b, "## Issues by Category\n\n")

	// Group results
	byCat := map[string][]score.ScoredRule{}
	for _, r := range m.Results {
		byCat[r.Category] = append(byCat[r.Category], r)
	}

	for cat, list := range byCat {
		fmt.Fprintf(&b, "### %s\n\n", strings.Title(cat))
		if len(list) == 0 {
			fmt.Fprintf(&b, "(none)\n\n")
			continue
		}
		for _, r := range list {
			msg := r.Message
			if r.File != "" {
				msg += fmt.Sprintf(" *(File: %s)*", r.File)
			}
			fmt.Fprintf(&b, "- **[%s] %s** — %s\n", r.Level, r.ID, msg)
		}
		fmt.Fprintf(&b, "\n")
	}

	// Fixable
	if len(m.Fixable) > 0 {
		fmt.Fprintf(&b, "## Fixable Issues\n\n")
		for _, f := range m.Fixable {
			fmt.Fprintf(
				&b,
				"- **%s** → %s *(target: %s)*\n",
				f.RuleID, f.Message, f.FilePath,
			)
		}
		fmt.Fprintf(&b, "\n")
	}

	return b.String()
}

//
// CI exit code logic
//

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

func unwrapResults(scored []score.ScoredRule) []rules.Result {
	out := make([]rules.Result, 0, len(scored))
	for _, r := range scored {
		out = append(out, r.Result)
	}
	return out
}
