package rules

import "github.com/jbakchr/hewd/internal/scan"

func init() {
    RegisterRule(RuleReadmeMissing)
    RegisterRule(RuleLicenseMissing)
    RegisterRule(RuleContributingMissing)
}

func RuleReadmeMissing(s *scan.Summary) []Result {
    if s.Documentation["README.md"] {
        return nil
    }
    return []Result{{
        ID:      "DOC_README_MISSING",
        Level:   Error,
        Message: "README.md is missing.",
    }}
}

func RuleLicenseMissing(s *scan.Summary) []Result {
    if s.Documentation["LICENSE"] {
        return nil
    }
    return []Result{{
        ID:      "DOC_LICENSE_MISSING",
        Level:   Warn,
        Message: "LICENSE file not found.",
    }}
}

func RuleContributingMissing(s *scan.Summary) []Result {
    if s.Documentation["CONTRIBUTING.md"] {
        return nil
    }
    return []Result{{
        ID:      "DOC_CONTRIB_MISSING",
        Level:   Info,
        Message: "CONTRIBUTING.md is recommended for open‑source projects.",
    }}
}