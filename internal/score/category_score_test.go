package score_test

import (
	"testing"

	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/score"
)

func TestScoreByCategory(t *testing.T) {
	results := []score.ScoredRule{
		{Result: rules.Result{ID: "A", Level: rules.Error}, Category: "documentation"},
		{Result: rules.Result{ID: "B", Level: rules.Warn}, Category: "config"},
		{Result: rules.Result{ID: "C", Level: rules.Info}, Category: "structure"},
	}

	cfg := &config.Config{Weights: map[string]int{}}
	scores := score.ScoreByCategory(results, cfg)

	if scores.Documentation != 90 {
		t.Errorf("expected doc score 90, got %d", scores.Documentation)
	}

	if scores.Config != 95 {
		t.Errorf("expected config score 95, got %d", scores.Config)
	}

	if scores.Structure != 99 {
		t.Errorf("expected structure score 99, got %d", scores.Structure)
	}

	if scores.Overall != (90+95+99)/3 {
		t.Errorf("unexpected overall score")
	}
}
