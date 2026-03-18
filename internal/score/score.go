package score

import (
	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/rules"
)

// ScoredRule augments rules.Result with a category field for structured output.
type ScoredRule struct {
	rules.Result
	Category string `json:"category" yaml:"category"`
}

// ScoreResult wraps the overall score and rule results.
type ScoreResult struct {
	Score   int          `json:"score" yaml:"score"`
	Results []ScoredRule `json:"results" yaml:"results"`
}

// Score computes a health score between 0 and 100 based on rule severities.
// Uses config.Weights to override severity levels when defined.
func Score(results []rules.Result, cfg *config.Config) int {
	score := 100

	for _, r := range results {
		level := r.Level

		// Apply severity override if present
		if cfg != nil && cfg.Weights != nil {
			if override, ok := cfg.Weights[r.ID]; ok {
				level = rules.LevelFromInt(override)
			}
		}

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
		score = 0
	}
	if score > 100 {
		score = 100
	}

	return score
}
