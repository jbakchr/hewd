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

func newDoctorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Run diagnostic checks on the project",
		Long: `Run validation rules against the scanned project structure.
Respects settings from .hewd/config.yaml. Produces structured output and
supports CI-friendly exit codes with --fail-on.`,
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

			results := rules.RunAll(summary, cfg)
			combined := score.ScoreResult{
				Score:   score.Score(results, cfg),
				Results: results,
			}

			jsonOut, _ := cmd.Flags().GetBool("json")
			yamlOut, _ := cmd.Flags().GetBool("yaml")
			pretty, _ := cmd.Flags().GetBool("pretty")
			failOnStr, _ := cmd.Flags().GetString("fail-on")
			showScore, _ := cmd.Flags().GetBool("score")

			if jsonOut && yamlOut {
				return fmt.Errorf("cannot use --json and --yaml together")
			}

			// ----- JSON -----
			if jsonOut {
				var data []byte
				if pretty {
					data, _ = json.MarshalIndent(combined, "", "  ")
				} else {
					data, _ = json.Marshal(combined)
				}
				fmt.Println(string(data))
				return evaluateDoctorExitCode(results, failOnStr)
			}

			// ----- YAML -----
			if yamlOut {
				data, _ := yaml.Marshal(combined)
				fmt.Println(string(data))
				return evaluateDoctorExitCode(results, failOnStr)
			}

			// ----- Pretty Text -----
			if showScore {
				fmt.Printf("Project Health Score: %d/100\n\n", combined.Score)
			}

			if len(results) == 0 {
				fmt.Println("No issues found. Project looks healthy!")
			} else {
				fmt.Println("Doctor Results:")
				for _, r := range results {
					fmt.Printf("[%s] %s: %s\n", r.Level, r.ID, r.Message)
					if r.File != "" {
						fmt.Printf("  File: %s\n", r.File)
					}
				}
			}

			return evaluateDoctorExitCode(results, failOnStr)
		},
	}

	cmd.Flags().Bool("json", false, "Output results in JSON format")
	cmd.Flags().Bool("yaml", false, "Output results in YAML format")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON output")
	cmd.Flags().String("fail-on", "error", "Fail on this severity or above (info|warn|error)")
	cmd.Flags().Bool("score", false, "Show project maturity score")

	return cmd
}

func evaluateDoctorExitCode(results []rules.Result, failOnStr string) error {

	highest := rules.Info
	for _, r := range results {
		if rules.SeverityRank(r.Level) > rules.SeverityRank(highest) {
			highest = r.Level
		}
	}

	var failOn rules.Level
	switch strings.ToLower(failOnStr) {
	case "info":
		failOn = rules.Info
	case "warn":
		failOn = rules.Warn
	case "error":
		failOn = rules.Error
	default:
		return fmt.Errorf("invalid --fail-on value: %s", failOnStr)
	}

	if rules.SeverityRank(highest) >= rules.SeverityRank(failOn) {
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
