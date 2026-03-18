package rules

import (
	"strings"
)

type Level string

const (
	Info  Level = "info"
	Warn  Level = "warn"
	Error Level = "error"
)

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

type Result struct {
	ID      string `json:"id" yaml:"id"`
	Level   Level  `json:"level" yaml:"level"`
	Message string `json:"message" yaml:"message"`
	File    string `json:"file,omitempty" yaml:"file,omitempty"`
}

type Rule func(summary interface{}) []Result

type RegisteredRule struct {
	ID       string
	Category string
	Func     Rule
}

var allRules []RegisteredRule

func RegisterRule(id string, category string, r Rule) {
	allRules = append(allRules, RegisteredRule{
		ID:       id,
		Category: category,
		Func:     r,
	})
}

func CategoryForRule(id string) string {
	for _, r := range allRules {
		if r.ID == id {
			return r.Category
		}
	}
	return "unknown"
}

func AllRules() []RegisteredRule {
	return allRules
}
