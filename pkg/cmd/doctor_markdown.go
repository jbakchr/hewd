package cmd

import (
	"fmt"
	"strings"

	"github.com/jbakchr/hewd/internal/api"
	"github.com/jbakchr/hewd/internal/rules"
)

// renderMarkdown produces a GitHub‑friendly Markdown doctor report.
// It intentionally mirrors the JSON/YAML machine output but with
// a human‑readable structure suitable for PR comments.
func renderMarkdown(m api.MachineOutput) string {
	var b strings.Builder

	// Header
	b.WriteString("# 📋 Hewd Doctor Report\n\n")

	// Score summary
	b.WriteString("## 🧮 Score Summary\n")
	b.WriteString(fmt.Sprintf("- **Overall Score:** %d/100\n", m.Score))
	b.WriteString(fmt.Sprintf("- **Documentation:** %d\n", m.CategoryScores.Documentation))
	b.WriteString(fmt.Sprintf("- **Config:** %d\n", m.CategoryScores.Config))
	b.WriteString(fmt.Sprintf("- **Structure:** %d\n\n", m.CategoryScores.Structure))

	// Issues grouped by category
	b.WriteString("## 🧩 Issues by Category\n\n")

	grouped := map[string][]rules.Result{}
	for _, sr := range m.Results {
		cat := rules.CategoryForRule(sr.ID)
		grouped[cat] = append(grouped[cat], sr.Result)
	}

	orderedCats := []string{"documentation", "config", "structure"}

	for _, cat := range orderedCats {
		issues := grouped[cat]

		b.WriteString(fmt.Sprintf("### %s\n", strings.Title(cat)))
		if len(issues) == 0 {
			b.WriteString("_No issues found._\n\n")
			continue
		}

		for _, r := range issues {
			icon, _ := rules.SeverityVisual(r.Level)

			if r.File != "" {
				b.WriteString(fmt.Sprintf(
					"- %s **%s** — %s _(file: %s)_\n",
					icon, r.ID, r.Message, r.File,
				))
			} else {
				b.WriteString(fmt.Sprintf(
					"- %s **%s** — %s\n",
					icon, r.ID, r.Message,
				))
			}
		}
		b.WriteString("\n")
	}

	// Fixable items
	if len(m.Fixable) > 0 {
		b.WriteString("## 🔧 Fixable Items\n")
		for _, f := range m.Fixable {
			b.WriteString(fmt.Sprintf(
				"- **%s** — %s (file: %s)\n",
				f.RuleID, f.Message, f.FilePath,
			))
		}
		b.WriteString("\n")
	}

	return b.String()
}
