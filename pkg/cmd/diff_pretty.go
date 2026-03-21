package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jbakchr/hewd/internal/cliutils"
	"github.com/jbakchr/hewd/internal/diff"
	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/score"
)

func writePrettyDiff(d diff.DiffResult) error {

	// ===== OVERALL SCORE =====
	fmt.Printf("%s===== OVERALL SCORE =====%s\n", cliutils.CyanBold, cliutils.Reset)
	fmt.Printf("  Change: %s\n\n", trendVisual(d.ScoreDelta))

	// ===== CATEGORY SCORE DELTAS =====
	fmt.Printf("%s===== CATEGORY SCORE DELTAS =====%s\n", cliutils.CyanBold, cliutils.Reset)
	if len(d.CategoryDeltas) == 0 {
		fmt.Println("  (none)")
	} else {
		cats := make([]string, 0, len(d.CategoryDeltas))
		for k := range d.CategoryDeltas {
			cats = append(cats, k)
		}
		sort.Strings(cats)

		for _, cat := range cats {
			delta := d.CategoryDeltas[cat]
			fmt.Printf("  %s: %s\n",
				uppercase(cat),
				trendVisual(delta),
			)
		}
		fmt.Println()
	}

	// ===== NEW ISSUES =====
	fmt.Printf("%s===== NEW ISSUES =====%s\n", cliutils.CyanBold, cliutils.Reset)
	if len(d.NewIssues) == 0 {
		fmt.Println("  (none)")
	} else {
		printIssues(d.NewIssues)
		fmt.Println()
	}

	// ===== RESOLVED ISSUES =====
	fmt.Printf("%s===== RESOLVED ISSUES =====%s\n", cliutils.CyanBold, cliutils.Reset)
	if len(d.ResolvedIssues) == 0 {
		fmt.Println("  (none)")
	} else {
		printIssues(d.ResolvedIssues)
		fmt.Println()
	}

	return nil
}

// Shared issue printer for new/resolved issues
func printIssues(list []score.ScoredRule) {
	for _, r := range list {
		icon, color := rules.SeverityVisual(r.Level)
		cat := rules.CategoryForRule(r.ID)

		fileSuffix := ""
		if r.File != "" {
			fileSuffix = fmt.Sprintf(" (%s)", r.File)
		}

		fmt.Printf("  %s%s%s  [%s] %s — %s%s\n",
			color, icon, cliutils.Reset,
			cat,
			r.ID,
			r.Message,
			fileSuffix,
		)
	}
}

// Colorized trend visual
func trendVisual(n int) string {
	switch {
	case n > 0:
		return fmt.Sprintf("%s+%d ↑%s", cliutils.Green, n, cliutils.Reset)
	case n < 0:
		return fmt.Sprintf("%s%d ↓%s", cliutils.Red, n, cliutils.Reset)
	default:
		return fmt.Sprintf("%s0%s", cliutils.WhiteBold, cliutils.Reset)
	}
}

// Uppercase first letter
func uppercase(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
