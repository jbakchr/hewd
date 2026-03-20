package diff

import (
	"sort"

	"github.com/jbakchr/hewd/internal/score"
)

var severityRank = map[string]int{
	"error": 1,
	"warn":  2,
	"info":  3,
}

func SortIssues(issues []score.ScoredRule) []score.ScoredRule {
	sorted := make([]score.ScoredRule, len(issues))
	copy(sorted, issues)

	sort.Slice(sorted, func(i, j int) bool {
		a := sorted[i]
		b := sorted[j]

		// Convert rule levels into raw strings for map indexing.
		aLevel := string(a.Level)
		bLevel := string(b.Level)

		// 1. Severity
		if severityRank[aLevel] != severityRank[bLevel] {
			return severityRank[aLevel] < severityRank[bLevel]
		}

		// 2. Category
		if a.Category != b.Category {
			return a.Category < b.Category
		}

		// 3. Rule ID
		if a.ID != b.ID {
			return a.ID < b.ID
		}

		// 4. Message
		return a.Message < b.Message
	})

	return sorted
}
