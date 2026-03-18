package score_test

import (
	"testing"

	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/score"
)

func TestScoreBasic(t *testing.T) {
	results := []rules.Result{
		{ID: "A", Level: rules.Error},
		{ID: "B", Level: rules.Warn},
		{ID: "C", Level: rules.Info},
	}

	s := score.Score(results, nil)

	if s != (100 - 10 - 5 - 1) {
		t.Fatalf("unexpected score: %d", s)
	}
}

func TestScoreWithOverrides(t *testing.T) {
	results := []rules.Result{
		{ID: "A", Level: rules.Warn}, // overridden below
	}

	cfg := &config.Config{
		Weights: map[string]int{
			"A": 3, // override to error
		},
	}

	s := score.Score(results, cfg)
	if s != 90 {
		t.Fatalf("expected score 90, got %d", s)
	}
}
