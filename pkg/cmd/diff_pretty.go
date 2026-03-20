package cmd

import (
	"fmt"

	"github.com/jbakchr/hewd/internal/api"
	"github.com/jbakchr/hewd/internal/diff"
)

func writePrettyDiff(result diff.DiffResult, old api.MachineOutput, new api.MachineOutput) error {

	groupedNew := diff.GroupIssues(result.NewIssues)
	groupedResolved := diff.GroupIssues(result.ResolvedIssues)

	fmt.Println("\n===== OVERALL SCORE =====")
	fmt.Println("----------------------------------------")
	fmt.Printf("Old report score: %d\n", old.Score)
	fmt.Printf("New report score: %d\n", new.Score)
	fmt.Printf("Score change: %+d\n", result.ScoreDelta)

	fmt.Println("\n===== CATEGORY SCORES =====")
	fmt.Println("----------------------------------------")

	fmt.Printf("  %-14s %3d → %3d   (%+d)\n",
		"Documentation:",
		old.CategoryScores.Documentation,
		new.CategoryScores.Documentation,
		result.CategoryDeltas["documentation"],
	)

	fmt.Printf("  %-14s %3d → %3d   (%+d)\n",
		"Config:",
		old.CategoryScores.Config,
		new.CategoryScores.Config,
		result.CategoryDeltas["config"],
	)

	fmt.Printf("  %-14s %3d → %3d   (%+d)\n",
		"Structure:",
		old.CategoryScores.Structure,
		new.CategoryScores.Structure,
		result.CategoryDeltas["structure"],
	)

	fmt.Println("\n===== NEW ISSUES =====")
	fmt.Println("----------------------------------------")
	if len(result.NewIssues) == 0 {
		fmt.Println("(none)")
	} else {
		for category, issues := range groupedNew {
			fmt.Printf("\n## %s\n", category)
			for _, issue := range issues {
				fmt.Printf("  - %s (%s)\n", issue.ID, issue.Level)
			}
		}
	}

	fmt.Println("\n===== RESOLVED ISSUES =====")
	fmt.Println("----------------------------------------")
	if len(result.ResolvedIssues) == 0 {
		fmt.Println("(none)")
	} else {
		for category, issues := range groupedResolved {
			fmt.Printf("\n## %s\n", category)
			for _, issue := range issues {
				fmt.Printf("  - %s (%s)\n", issue.ID, issue.Level)
			}
		}
	}

	return nil
}
