package cmd

import (
    "fmt"

    idoctor "github.com/jbakchr/hewd/internal/doctor"
)

func printDoctorResults(r idoctor.Result) {
    fmt.Println("Diagnostics:")

    for _, f := range r.Findings {
        status := "FAIL"
        if f.Passed {
            status = "OK"
        }

        fmt.Printf("  [%s] %s — %s\n", status, f.RuleID, f.Message)
    }

    fmt.Println("\nDoctor complete.")
}