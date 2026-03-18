package rules

import (
	"github.com/jbakchr/hewd/internal/scan"
)

// This file contains the *fundamental* documentation rules for hewd.
// Additional rules are provided in rules_documentation_extra.go.

// -----------------------------------------------------------------------------
// Rule Registration
// -----------------------------------------------------------------------------

func init() {
	RegisterRule("DOC_README_MISSING", RuleReadmeMissing)
	RegisterRule("DOC_LICENSE_MISSING", RuleLicenseMissing)
	RegisterRule("DOC_CONTRIBUTING_MISSING", RuleContributingMissing)
}

// -----------------------------------------------------------------------------
// Rules
// -----------------------------------------------------------------------------

// RuleReadmeMissing checks if README.md exists.
func RuleReadmeMissing(s interface{}) []Result {
	summary := s.(*scan.Summary)

	if summary.Documentation["README.md"] {
		return nil
	}

	return []Result{{
		ID:      "DOC_README_MISSING",
		Level:   Error,
		Message: "README.md is missing.",
	}}
}

// RuleLicenseMissing checks if LICENSE exists.
func RuleLicenseMissing(s interface{}) []Result {
	summary := s.(*scan.Summary)

	if summary.Documentation["LICENSE"] {
		return nil
	}

	return []Result{{
		ID:      "DOC_LICENSE_MISSING",
		Level:   Warn,
		Message: "LICENSE file is missing.",
	}}
}

// RuleContributingMissing checks if CONTRIBUTING.md exists.
func RuleContributingMissing(s interface{}) []Result {
	summary := s.(*scan.Summary)

	if summary.Documentation["CONTRIBUTING.md"] {
		return nil
	}

	return []Result{{
		ID:      "DOC_CONTRIBUTING_MISSING",
		Level:   Info,
		Message: "CONTRIBUTING.md is recommended for open-source projects, but is missing.",
	}}
}
