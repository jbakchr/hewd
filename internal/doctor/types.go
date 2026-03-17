package doctor

import "github.com/jbakchr/hewd/internal/scan"

// Rule represents a single diagnostic rule.
type Rule struct {
    ID          string
    Description string
    Check       func(*scan.Summary) (passed bool, message string)
}

// Finding is the result of running one rule.
type Finding struct {
    RuleID      string `json:"rule_id" yaml:"rule_id"`
    Passed      bool   `json:"passed" yaml:"passed"`
    Message     string `json:"message" yaml:"message"`
}

// Result aggregates all rule findings.
type Result struct {
    Findings []Finding `json:"findings" yaml:"findings"`
}
