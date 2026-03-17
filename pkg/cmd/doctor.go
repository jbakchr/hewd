package cmd

import (
    "encoding/json"
    "fmt"
    "os"
    "strings"

    "github.com/spf13/cobra"
    "gopkg.in/yaml.v3"

    "github.com/jbakchr/hewd/internal/scan"
    "github.com/jbakchr/hewd/internal/rules"
)

func newDoctorCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "doctor",
        Short: "Run diagnostic checks on the project",
        RunE: func(cmd *cobra.Command, args []string) error {

            cwd, err := os.Getwd()
            if err != nil {
                return err
            }

            summary, err := scan.ScanDirectory(cwd)
            if err != nil {
                return err
            }

            results := rules.RunAll(summary)

            jsonOut, _ := cmd.Flags().GetBool("json")
            yamlOut, _ := cmd.Flags().GetBool("yaml")
            pretty, _ := cmd.Flags().GetBool("pretty")
            failOnStr, _ := cmd.Flags().GetString("fail-on")

            if jsonOut && yamlOut {
                return fmt.Errorf("cannot use --json and --yaml together")
            }

            // JSON output
            if jsonOut {
                var data []byte
                if pretty {
                    data, _ = json.MarshalIndent(results, "", "  ")
                } else {
                    data, _ = json.Marshal(results)
                }
                fmt.Println(string(data))
            } else if yamlOut {
                data, _ := yaml.Marshal(results)
                fmt.Println(string(data))
            } else {
                // Human-readable output
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
            }

            // Determine highest severity detected
            highest := rules.Info
            for _, r := range results {
                if rules.SeverityRank(r.Level) > rules.SeverityRank(highest) {
                    highest = r.Level
                }
            }

            // Parse fail-on
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

            // Evaluate whether to fail
            if rules.SeverityRank(highest) >= rules.SeverityRank(failOn) {
                // exit code mapping
                if highest == rules.Error {
                    os.Exit(2)
                }
                os.Exit(1)
            }

            return nil
        },
    }

    cmd.Flags().Bool("json", false, "Output results in JSON format")
    cmd.Flags().Bool("yaml", false, "Output results in YAML format")
    cmd.Flags().Bool("pretty", false, "Pretty-print JSON output")
    cmd.Flags().String("fail-on", "error", "Fail the command if this level or above occurs (info|warn|error)")

    return cmd
}