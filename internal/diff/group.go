package diff

import "github.com/jbakchr/hewd/internal/score"

type GroupedIssues map[string][]score.ScoredRule

func GroupIssues(issues []score.ScoredRule) GroupedIssues {
	groups := make(GroupedIssues)

	for _, issue := range issues {
		cat := issue.Category
		groups[cat] = append(groups[cat], issue)
	}

	// Sort issues in each category
	for cat := range groups {
		groups[cat] = SortIssues(groups[cat])
	}

	return groups
}
