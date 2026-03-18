package rules

import (
    "strings"
)

// Level represents the severity of a rule result.
type Level string

const (
    Info  Level = "info"
    Warn  Level = "warn"
    Error Level = "error"
)

// LevelFromInt converts an integer severity configured in .hewd/config.yaml
// into a Level enum. You can expand this later if needed.
func LevelFromInt(i int) Level {
    switch i {
    case 3:
        return Error
    case 2:
        return Warn
    case 1:
        return Info
    }
    return Info
}

// SeverityRank assigns a numeric weight to a severity level so we can
// compare their importance (for fail-on flags and rule evaluation).
func SeverityRank(l Level) int {
    switch strings.ToLower(string(l)) {
    case "info":
        return 1
    case "warn":
        return 2
    case "error":
        return 3
    default:
        return 0
    }
}

// Result represents a single rule violation or informational diagnostic.
type Result struct {
    ID      string `json:"id" yaml:"id"`
    Level   Level  `json:"level" yaml:"level"`
    Message string `json:"message" yaml:"message"`
    File    string `json:"file,omitempty" yaml:"file,omitempty"`
}

// Rule represents a rule function, which takes a scan Summary and returns results.
type Rule func(summary interface{}) []Result

// RegisteredRule associates a rule with its unique ID.
type RegisteredRule struct {
    ID   string
    Func Rule
}

// allRules stores the list of registered rules.
// These rules are appended via RegisterRule().
var allRules []RegisteredRule

// RegisterRule registers a rule with its explicit rule ID.
// Example: RegisterRule("DOC_README_MISSING", RuleReadmeMissing)
func RegisterRule(id string, r Rule) {
    allRules = append(allRules, RegisteredRule{
        ID:   id,
        Func: r,
    })
}

// AllRules returns all rules registered so far.
// Useful for testing or introspection.
func AllRules() []RegisteredRule {
    return allRules
}