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
)

func newDoctorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Run diagnostic checks on the project",
		Long: `Run validation rules against the scanned project structure.
Respects settings from .hewd/config.yaml. Produces structured output and
supports CI-friendly exit codes with --fail-on.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			// Load project root
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			// Load config (missing config is not an error)
			cfg, _ := config.Load(cwd)

			// Perform directory scan
			summary, err := scan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			// Run rules with config
			results := rules.RunAll(summary, cfg)

			// Read flags
			jsonOut, _ := cmd.Flags().GetBool("json")
			yamlOut, _ := cmd.Flags().GetBool("yaml")
			pretty, _ := cmd.Flags().GetBool("pretty")
			failOnStr, _ := cmd.Flags().GetString("fail-on")

			// Validate incompatible flags
			if jsonOut && yamlOut {
				return fmt.Errorf("cannot use --json and --yaml together")
			}

			// -----------------------------------------------------------
			// Output: JSON
			// -----------------------------------------------------------
			if jsonOut {
				var data []byte
				if pretty {
					data, _ = json.MarshalIndent(results, "", "  ")
				} else {
					data, _ = json.Marshal(results)
				}
				fmt.Println(string(data))
				return evaluateDoctorExitCode(results, failOnStr)
			}

			// -----------------------------------------------------------
			// Output: YAML
			// -----------------------------------------------------------
			if yamlOut {
				data, _ := yaml.Marshal(results)
				fmt.Println(string(data))
				return evaluateDoctorExitCode(results, failOnStr)
			}

			// -----------------------------------------------------------
			// Output: human-readable text
			// -----------------------------------------------------------
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

	return cmd
}

// -----------------------------------------------------------------------------
// Exit code evaluation
// -----------------------------------------------------------------------------

func evaluateDoctorExitCode(results []rules.Result, failOnStr string) error {

	// Determine highest severity present
	highest := rules.Info
	for _, r := range results {
		if rules.SeverityRank(r.Level) > rules.SeverityRank(highest) {
			highest = r.Level
		}
	}

	// Parse fail-on flag
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

	// Compare severities
	if rules.SeverityRank(highest) >= rules.SeverityRank(failOn) {
		// Map severity to distinct exit codes
		switch highest {
		case rules.Error:
			os.Exit(2)
		case rules.Warn:
			os.Exit(1)
		case rules.Info:
			os.Exit(1) // info only fails if fail-on=info
		}
	}

	// OK
	return nil
}
