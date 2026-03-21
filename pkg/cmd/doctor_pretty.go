package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jbakchr/hewd/internal/cliutils"
	"github.com/jbakchr/hewd/internal/rules"
)

func printDoctorPretty(results []rules.Result) {

	// Group results by category USING Option C
	grouped := map[string][]rules.Result{}
	for _, r := range results {
		cat := rules.CategoryForRule(r.ID)
		grouped[cat] = append(grouped[cat], r)
	}

	// Stable category display ordering
	orderedCats := []string{"documentation", "config", "structure"}
	// Add any custom categories registered later
	for cat := range grouped {
		if !contains(orderedCats, cat) {
			orderedCats = append(orderedCats, cat)
		}
	}

	for _, cat := range orderedCats {

		fmt.Printf("%s===== %s ISSUES =====%s\n",
			cliutils.CyanBold,
			strings.ToUpper(cat),
			cliutils.Reset)

		issues := grouped[cat]

		if len(issues) == 0 {
			fmt.Println("  (none)")
			continue
		}

		// Sort: errors first, warnings next, then info
		sort.Slice(issues, func(i, j int) bool {
			return rules.SeverityRank(issues[i].Level) > rules.SeverityRank(issues[j].Level)
		})

		for _, r := range issues {
			icon, color := cliutils.SeverityVisual(r.Level)

			fileSuffix := ""
			if r.File != "" {
				fileSuffix = fmt.Sprintf(" (%s)", r.File)
			}

			fmt.Printf("  %s%s%s  %s%s%s — %s%s\n",
				color, icon, cliutils.Reset,
				cliutils.Bold, r.ID, cliutils.Reset,
				r.Message,
				fileSuffix,
			)
		}

		fmt.Println()
	}
}

func contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
