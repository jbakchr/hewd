package rules

import "github.com/jbakchr/hewd/internal/cliutils"

func SeverityVisual(level Level) (string, string) {
	switch level {
	case Error:
		return "❌", cliutils.RedBold
	case Warn:
		return "⚠️ ", cliutils.Yellow
	default:
		return "ℹ️ ", cliutils.Blue
	}
}
