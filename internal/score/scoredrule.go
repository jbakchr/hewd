package score

import (
	"github.com/jbakchr/hewd/internal/rules"
)

// NewScoredRules wraps plain rule results with category metadata.
func NewScoredRules(results []rules.Result) []ScoredRule {
	out := make([]ScoredRule, 0, len(results))

	for _, r := range results {
		out = append(out, ScoredRule{
			Result:   r,
			Category: rules.CategoryForRule(r.ID),
		})
	}

	return out
}
