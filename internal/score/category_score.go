package score

import (
	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/rules"
)

// CategoryScores holds individual category scores plus the overall score.
type CategoryScores struct {
	Documentation int `json:"documentation" yaml:"documentation"`
	Config        int `json:"config" yaml:"config"`
	Structure     int `json:"structure" yaml:"structure"`
	Overall       int `json:"overall" yaml:"overall"`
}

func ScoreByCategory(results []ScoredRule, cfg *config.Config) CategoryScores {
	scores := CategoryScores{
		Documentation: 100,
		Config:        100,
		Structure:     100,
	}

	// Deduct category-specific points
	for _, r := range results {
		level := r.Level

		// severity override
		if cfg != nil && cfg.Weights != nil {
			if override, ok := cfg.Weights[r.ID]; ok {
				level = rules.LevelFromInt(override)
			}
		}

		switch r.Category {
		case "documentation":
			scores.Documentation -= penalty(level)
		case "config":
			scores.Config -= penalty(level)
		case "structure":
			scores.Structure -= penalty(level)
		}
	}

	// Clamp each category
	if scores.Documentation < 0 {
		scores.Documentation = 0
	}
	if scores.Config < 0 {
		scores.Config = 0
	}
	if scores.Structure < 0 {
		scores.Structure = 0
	}

	// Overall score = average of the three
	scores.Overall = (scores.Documentation + scores.Config + scores.Structure) / 3
	return scores
}

func penalty(level rules.Level) int {
	switch level {
	case rules.Error:
		return 10
	case rules.Warn:
		return 5
	case rules.Info:
		return 1
	default:
		return 0
	}
}
