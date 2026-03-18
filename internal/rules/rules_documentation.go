package rules

import "github.com/jbakchr/hewd/internal/scan"

func init() {
	RegisterRule("DOC_README_MISSING", "documentation", RuleReadmeMissing)
	RegisterRule("DOC_LICENSE_MISSING", "documentation", RuleLicenseMissing)
	RegisterRule("DOC_CONTRIBUTING_MISSING", "documentation", RuleContributingMissing)
}

func RuleReadmeMissing(s interface{}) []Result {
	sum := s.(*scan.Summary)

	if sum.Documentation["README.md"] {
		return nil
	}

	return []Result{{
		ID:      "DOC_README_MISSING",
		Level:   Error,
		Message: "README.md is missing.",
	}}
}

func RuleLicenseMissing(s interface{}) []Result {
	sum := s.(*scan.Summary)

	if sum.Documentation["LICENSE"] {
		return nil
	}

	return []Result{{
		ID:      "DOC_LICENSE_MISSING",
		Level:   Warn,
		Message: "LICENSE file is missing.",
	}}
}

func RuleContributingMissing(s interface{}) []Result {
	sum := s.(*scan.Summary)

	if sum.Documentation["CONTRIBUTING.md"] {
		return nil
	}

	return []Result{{
		ID:      "DOC_CONTRIBUTING_MISSING",
		Level:   Info,
		Message: "CONTRIBUTING.md is recommended for open‑source projects.",
	}}
}
