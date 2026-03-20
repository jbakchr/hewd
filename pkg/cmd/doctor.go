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
		Short: "Run full diagnostics and compute documentation, config, and structure scores.",
		Long: `hewd doctor runs the full diagnostic engine on the current repository. It
evaluates documentation, configuration, and structure using a curated set of
rules, each with a severity level (info, warn, error).

The doctor command produces:

  • Category scores (documentation, config, structure)
  • An overall project health score
  • Detailed issue reports with severity metadata
  • A list of fixable problems
  • Optional Markdown, JSON, or YAML output

Markdown output is ideal for pull request comments, while JSON and YAML are
well‑suited for CI pipelines and automated quality gates.

Use 'hewd doctor' regularly to enforce documentation standards, detect
regressions, and maintain consistent quality across repositories.`,
		Example: `
  # Run full diagnostics using pretty output (default)
  hewd doctor

  # Output Markdown report (ideal for PR comments)
  hewd doctor --md > health.md

  # Output JSON for CI pipelines or dashboards
  hewd doctor --json > doctor.json

  # Output YAML
  hewd doctor --yaml

  # Only evaluate documentation-related rules
  hewd doctor --only documentation

  # Exclude config-related checks
  hewd doctor --except config

  # Fail CI if any warning-level issues occur
  hewd doctor --fail-on=warn
`,
		RunE: func(cmd *cobra.Command, args []string) error {

			// ----- Directory -----
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not get working directory: %w", err)
			}

			// Config (optional)
			cfg, _ := config.Load(cwd)

			// Repo summary
			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			// ----- Flags -----
			onlyCats, _ := cmd.Flags().GetStringSlice("only")
			exceptCats, _ := cmd.Flags().GetStringSlice("except")

			jsonOut, _ := cmd.Flags().GetBool("json")
			yamlOut, _ := cmd.Flags().GetBool("yaml")
			markdownOut, _ := cmd.Flags().GetBool("md")
			prettyJSON, _ := cmd.Flags().GetBool("pretty")

			showScore, _ := cmd.Flags().GetBool("score")
			showCategoryScore, _ := cmd.Flags().GetBool("category-score")

			failOnStr, _ := cmd.Flags().GetString("fail-on")

			// ----- Flag Conflicts -----
			if (jsonOut && yamlOut) ||
				(jsonOut && markdownOut) ||
				(yamlOut && markdownOut) {
				return fmt.Errorf("cannot combine --json, --yaml, or --md")
			}

			if yamlOut && prettyJSON {
				return fmt.Errorf("cannot combine --yaml and --pretty (pretty applies only to JSON)")
			}

			if markdownOut && prettyJSON {
				return fmt.Errorf("cannot combine --md and --pretty (pretty applies only to JSON)")
			}

			// These flags only apply to pretty output
			if (showScore || showCategoryScore) &&
				(jsonOut || yamlOut || markdownOut) {
				return fmt.Errorf("--score and --category-score can only be used with pretty terminal output")
			}

			// ----- Run Rules -----
			results := rules.RunAll(summary, cfg, onlyCats, exceptCats)

			// Score wrapping
			scoredRules := score.NewScoredRules(results)

			categoryScores := score.ScoreByCategory(scoredRules, cfg)
			overallScore := score.Score(results, cfg)

			// Fixable detection
			rawFixes := fix.DetectFixes(results, cwd)
			var fixables []api.FixableItem
			for _, f := range rawFixes {
				fixables = append(fixables, api.FixableItem{
					RuleID:   f.RuleID,
					Message:  f.Message,
					FilePath: f.FilePath,
				})
			}

			// Machine-readable object
			machine := api.MachineOutput{
				SchemaVersion:  1,
				HewdVersion:    version.Version,
				GeneratedAt:    time.Now(),
				Score:          overallScore,
				CategoryScores: categoryScores,
				Results:        scoredRules,
				Fixable:        fixables,
			}

			// ----- Markdown -----
			if markdownOut {
				md := renderMarkdown(machine)
				fmt.Println(md)
				return evaluateDoctorExitCode(results, failOnStr)
			}

			// ----- JSON -----
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

			// ----- YAML -----
			if yamlOut {
				data, _ := yaml.Marshal(machine)
				fmt.Println(string(data))
				return evaluateDoctorExitCode(results, failOnStr)
			}

			// ----- Pretty Terminal Output -----
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

	// ----- Flags -----
	cmd.Flags().StringSlice("only", []string{}, "Only run rules from specific categories (comma-separated)")
	cmd.Flags().StringSlice("except", []string{}, "Skip rules from specific categories (comma-separated)")

	cmd.Flags().Bool("json", false, "Output the diagnostic report in JSON format. Use --pretty for indented JSON.")
	cmd.Flags().Bool("yaml", false, "Output the diagnostic report in YAML format.")
	cmd.Flags().Bool("md", false, "Output the diagnostic report in Markdown format.")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON output.")

	cmd.Flags().String("fail-on", "error", "Fail if a rule of this severity or higher occurs (info|warn|error)")
	cmd.Flags().Bool("score", false, "Print only the overall score (terminal output only)")
	cmd.Flags().Bool("category-score", false, "Print only category scores (terminal output only)")

	return cmd
}

//
// Markdown Output
//

func renderMarkdown(m api.MachineOutput) string {
	var b strings.Builder

	fmt.Fprintf(&b, "# hewd Report\n\n")

	fmt.Fprintf(&b, "Generated by **hewd v%s** on %s\n\n",
		m.HewdVersion, m.GeneratedAt.Format("2006-01-02 15:04:05 MST"))

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
// CI Exit Code Logic
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
