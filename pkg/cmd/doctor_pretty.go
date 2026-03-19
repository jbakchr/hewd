package cmd

import (
	"fmt"
	"strings"

	"github.com/jbakchr/hewd/internal/rules"
)

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
