package doctor

import "github.com/jbakchr/hewd/internal/scan"

// RunAll executes all built-in rules.
func RunAll(s *scan.Summary) Result {
    results := Result{}

    for _, rule := range Rules {
        passed, message := rule.Check(s)

        results.Findings = append(results.Findings, Finding{
            RuleID:  rule.ID,
            Passed:  passed,
            Message: message,
        })
    }

    return results
}