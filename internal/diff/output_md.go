package diff

import (
	"fmt"
	"strings"

	"github.com/jbakchr/hewd/internal/api"
)

func emojiForDelta(delta int) string {
	switch {
	case delta > 0:
		return "🟩"
	case delta < 0:
		return "🟥"
	default:
		return "⬜"
	}
}

func WriteMarkdown(result DiffResult, old api.MachineOutput, new api.MachineOutput) string {
	var b strings.Builder

	// -----------------------------------------------------
	// HEADER
	// -----------------------------------------------------
	b.WriteString("# 📊 Hewd Diff Report\n\n")
	b.WriteString("---\n\n")

	// -----------------------------------------------------
	// SCORE SUMMARY
	// -----------------------------------------------------
	b.WriteString("## 📈 Score Summary\n\n")

	b.WriteString("| Metric | Old | New | Δ |\n")
	b.WriteString("|--------|-----|-----|----|\n")

	// Overall score row
	b.WriteString(fmt.Sprintf(
		"| Overall Score | **%d** | **%d** | %s **%+d** |\n",
		old.Score, new.Score, emojiForDelta(result.ScoreDelta), result.ScoreDelta,
	))

	// Documentation
	docDelta := result.CategoryDeltas["documentation"]
	b.WriteString(fmt.Sprintf(
		"| Documentation | **%d** | **%d** | %s **%+d** |\n",
		old.CategoryScores.Documentation,
		new.CategoryScores.Documentation,
		emojiForDelta(docDelta),
		docDelta,
	))

	// Config
	cfgDelta := result.CategoryDeltas["config"]
	b.WriteString(fmt.Sprintf(
		"| Config | **%d** | **%d** | %s **%+d** |\n",
		old.CategoryScores.Config,
		new.CategoryScores.Config,
		emojiForDelta(cfgDelta),
		cfgDelta,
	))

	// Structure
	structDelta := result.CategoryDeltas["structure"]
	b.WriteString(fmt.Sprintf(
		"| Structure | **%d** | **%d** | %s **%+d** |\n",
		old.CategoryScores.Structure,
		new.CategoryScores.Structure,
		emojiForDelta(structDelta),
		structDelta,
	))

	b.WriteString("\n---\n\n")

	// -----------------------------------------------------
	// NEW ISSUES
	// -----------------------------------------------------
	groupedNew := GroupIssues(result.NewIssues)

	b.WriteString("## 🆕 New Issues\n\n")
	if len(result.NewIssues) == 0 {
		b.WriteString("_No new issues! 🎉_\n\n")
	} else {
		for category, issues := range groupedNew {
			b.WriteString(fmt.Sprintf("### %s\n", category))
			for _, issue := range issues {
				b.WriteString(fmt.Sprintf(
					"- **%s** (%s) — %s\n",
					issue.ID, issue.Level, issue.Message,
				))
			}
			b.WriteString("\n")
		}
	}

	b.WriteString("\n---\n\n")

	// -----------------------------------------------------
	// RESOLVED ISSUES
	// -----------------------------------------------------
	groupedResolved := GroupIssues(result.ResolvedIssues)

	b.WriteString("## ✅ Resolved Issues\n\n")
	if len(result.ResolvedIssues) == 0 {
		b.WriteString("_No resolved issues._\n\n")
	} else {
		for category, issues := range groupedResolved {
			b.WriteString(fmt.Sprintf("### %s\n", category))
			for _, issue := range issues {
				b.WriteString(fmt.Sprintf(
					"- **%s** (%s) — %s\n",
					issue.ID, issue.Level, issue.Message,
				))
			}
			b.WriteString("\n")
		}
	}

	b.WriteString("\n---\n\n")

	return b.String()
}
