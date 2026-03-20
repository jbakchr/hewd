package diff

import (
	"fmt"
	"strings"

	"github.com/jbakchr/hewd/internal/api"
)

// WriteMarkdown writes a GitHub-friendly Markdown diff report to stdout.
func WriteMarkdown(result DiffResult, old api.MachineOutput, new api.MachineOutput) string {
	var b strings.Builder

	// -----------------------------------------------------
	// Header
	// -----------------------------------------------------
	b.WriteString("# Hewd Diff Report\n\n")

	// -----------------------------------------------------
	// Score Summary
	// -----------------------------------------------------
	b.WriteString("## Score Summary\n\n")

	b.WriteString("| Metric | Old | New | Δ |\n")
	b.WriteString("|--------|-----|-----|----|\n")

	b.WriteString(fmt.Sprintf(
		"| Overall Score | %d | %d | %+d |\n",
		old.Score, new.Score, result.ScoreDelta,
	))

	b.WriteString(fmt.Sprintf(
		"| Documentation | %d | %d | %+d |\n",
		old.CategoryScores.Documentation,
		new.CategoryScores.Documentation,
		result.CategoryDeltas["documentation"],
	))

	b.WriteString(fmt.Sprintf(
		"| Config | %d | %d | %+d |\n",
		old.CategoryScores.Config,
		new.CategoryScores.Config,
		result.CategoryDeltas["config"],
	))

	b.WriteString(fmt.Sprintf(
		"| Structure | %d | %d | %+d |\n",
		old.CategoryScores.Structure,
		new.CategoryScores.Structure,
		result.CategoryDeltas["structure"],
	))

	b.WriteString("\n")

	// -----------------------------------------------------
	// New Issues
	// -----------------------------------------------------
	groupedNew := GroupIssues(result.NewIssues)

	b.WriteString("## New Issues\n\n")
	if len(result.NewIssues) == 0 {
		b.WriteString("_None_\n\n")
	} else {
		for category, issues := range groupedNew {
			b.WriteString(fmt.Sprintf("### %s\n", category))
			for _, issue := range issues {
				b.WriteString(fmt.Sprintf("- **%s** (%s): %s\n", issue.ID, issue.Level, issue.Message))
			}
			b.WriteString("\n")
		}
	}

	// -----------------------------------------------------
	// Resolved Issues
	// -----------------------------------------------------
	groupedResolved := GroupIssues(result.ResolvedIssues)

	b.WriteString("## Resolved Issues\n\n")
	if len(result.ResolvedIssues) == 0 {
		b.WriteString("_None_\n\n")
	} else {
		for category, issues := range groupedResolved {
			b.WriteString(fmt.Sprintf("### %s\n", category))
			for _, issue := range issues {
				b.WriteString(fmt.Sprintf("- **%s** (%s): %s\n", issue.ID, issue.Level, issue.Message))
			}
			b.WriteString("\n")
		}
	}

	return b.String()
}
