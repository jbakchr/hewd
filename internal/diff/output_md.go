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

	b.WriteString("# 📊 Hewd Diff Report\n\n---\n\n")

	b.WriteString("## 📈 Score Summary\n\n")

	b.WriteString("| Metric         | Old  | New  | Δ     | Trend |\n")
	b.WriteString("|----------------|------|------|-------|--------|\n")

	// Helper for table row
	writeRow := func(name string, oldVal, newVal, delta int) {
		emoji := emojiForDelta(delta)
		arrow := "➡️"
		if delta > 0 {
			arrow = "⬆️"
		} else if delta < 0 {
			arrow = "⬇️"
		}
		b.WriteString(fmt.Sprintf(
			"| %-14s | %4d | %4d | %+5d | %s%s |\n",
			name, oldVal, newVal, delta, emoji, arrow,
		))
	}

	writeRow("Overall Score", old.Score, new.Score, result.ScoreDelta)
	writeRow("Documentation", old.CategoryScores.Documentation, new.CategoryScores.Documentation, result.CategoryDeltas["documentation"])
	writeRow("Config", old.CategoryScores.Config, new.CategoryScores.Config, result.CategoryDeltas["config"])
	writeRow("Structure", old.CategoryScores.Structure, new.CategoryScores.Structure, result.CategoryDeltas["structure"])

	b.WriteString("\n---\n\n")

	// New Issues
	groupedNew := GroupIssues(result.NewIssues)
	b.WriteString("## 🆕 New Issues\n\n")
	if len(result.NewIssues) == 0 {
		b.WriteString("_No new issues! 🎉_\n\n")
	} else {
		for category, issues := range groupedNew {
			b.WriteString(fmt.Sprintf("### %s\n", category))
			for _, issue := range issues {
				b.WriteString(fmt.Sprintf("- **%s** (%s) — %s\n", issue.ID, issue.Level, issue.Message))
			}
			b.WriteString("\n")
		}
	}

	b.WriteString("\n---\n\n")

	// Resolved Issues
	groupedResolved := GroupIssues(result.ResolvedIssues)
	b.WriteString("## ✅ Resolved Issues\n\n")
	if len(result.ResolvedIssues) == 0 {
		b.WriteString("_No resolved issues._\n\n")
	} else {
		for category, issues := range groupedResolved {
			b.WriteString(fmt.Sprintf("### %s\n", category))
			for _, issue := range issues {
				b.WriteString(fmt.Sprintf("- **%s** (%s) — %s\n", issue.ID, issue.Level, issue.Message))
			}
			b.WriteString("\n")
		}
	}

	b.WriteString("\n---\n\n")

	return b.String()
}
