package score

import (
	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/rules"
)

// ScoreResult wraps score + rule results (for JSON/YAML output)
type ScoreResult struct {
	Score   int            `json:"score" yaml:"score"`
	Results []rules.Result `json:"results" yaml:"results"`
}

// Score computes a 0–100 maturity score based on rule severities.
// Uses optional config weights to override severity levels.
func Score(results []rules.Result, cfg *config.Config) int {

	// Base score
	score := 100

	for _, r := range results {
		// Apply config severity overrides
		level := r.Level
		if cfg != nil && cfg.Weights != nil {
			if override, ok := cfg.Weights[r.ID]; ok {
				level = rules.LevelFromInt(override)
			}
		}

		// Subtract points based on severity
		switch level {
		case rules.Error:
			score -= 10
		case rules.Warn:
			score -= 5
		case rules.Info:
			score -= 1
		}
	}

	if score < 0 {
		return 0
	}
	if score > 100 {
		return 100
	}
	return score
}
