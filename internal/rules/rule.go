package rules

import "github.com/jbakchr/hewd/internal/scan"

type Level string

const (
	Info  Level = "info"
	Warn  Level = "warn"
	Error Level = "error"
)

type Result struct {
	ID      string `json:"id" yaml:"id"`
	Level   Level  `json:"level" yaml:"level"`
	Message string `json:"message" yaml:"message"`
	File    string `json:"file,omitempty" yaml:"file,omitempty"`
}

type Rule func(*scan.Summary) []Result

func SeverityRank(l Level) int {
	switch l {
	case Info:
		return 1
	case Warn:
		return 2
	case Error:
		return 3
	default:
		return 0
	}
}
