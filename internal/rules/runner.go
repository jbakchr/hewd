package rules

import "github.com/jbakchr/hewd/internal/scan"

var allRules []Rule

func RegisterRule(r Rule) {
	allRules = append(allRules, r)
}

func RunAll(summary interface {
	// we expect same interface as scan.Summary
}) []Result {

	results := []Result{}

	for _, rule := range allRules {
		out := rule(summary.(*scan.Summary))
		results = append(results, out...)
	}

	return results
}
