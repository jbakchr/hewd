package doctor

import "github.com/jbakchr/hewd/internal/scan"

func RunAll(s *scan.Summary) Result {
	result := Result{}

	total := 0
	max := 0

	for _, rule := range Rules {

		// Default weight = 1 if none is set
		w := rule.Weight
		if w <= 0 {
			w = 1
		}

		passed, message := rule.Check(s)

		result.Findings = append(result.Findings, Finding{
			RuleID:  rule.ID,
			Passed:  passed,
			Message: message,
		})

		// Count towards scoring
		max += w
		if passed {
			total += w
		}
	}

	result.Score = total
	result.MaxScore = max

	return result
}
