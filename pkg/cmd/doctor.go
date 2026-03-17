package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
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

            if len(results) == 0 {
                fmt.Println("No issues found. Project looks healthy!")
                return nil
            }

            fmt.Println("Doctor Results:")
            for _, r := range results {
                fmt.Printf("[%s] %s: %s\n", r.Level, r.ID, r.Message)
                if r.File != "" {
                    fmt.Printf("  File: %s\n", r.File)
                }
            }

            return nil
        },
    }

    return cmd
}