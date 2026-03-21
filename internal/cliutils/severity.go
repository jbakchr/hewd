package cliutils

import "github.com/jbakchr/hewd/internal/rules"

// SeverityVisual returns (icon, colorCode) for the given severity level.
func SeverityVisual(level rules.Level) (string, string) {
	switch level {
	case rules.Error:
		return "❌", RedBold
	case rules.Warn:
		return "⚠️ ", Yellow
	default:
		return "ℹ️ ", Blue
	}
}
