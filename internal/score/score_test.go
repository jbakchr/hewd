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
	expected := 100 - 10 - 5 - 1

	if s != expected {
		t.Fatalf("expected %d, got %d", expected, s)
	}
}

func TestScoreWithOverrides(t *testing.T) {
	results := []rules.Result{
		{ID: "A", Level: rules.Warn}, // will be overridden to Error
	}

	cfg := &config.Config{
		Weights: map[string]int{
			"A": 3, // Error
		},
	}

	s := score.Score(results, cfg)
	expected := 100 - 10

	if s != expected {
		t.Fatalf("expected %d, got %d", expected, s)
	}
}

func TestScoreBelowZero(t *testing.T) {
	results := []rules.Result{
		{ID: "A", Level: rules.Error},
		{ID: "B", Level: rules.Error},
		{ID: "C", Level: rules.Error},
		{ID: "D", Level: rules.Error},
		{ID: "E", Level: rules.Error},
		{ID: "F", Level: rules.Error},
		{ID: "G", Level: rules.Error},
		{ID: "H", Level: rules.Error},
		{ID: "I", Level: rules.Error},
		{ID: "J", Level: rules.Error},
		{ID: "K", Level: rules.Error},
	}

	s := score.Score(results, nil)
	if s != 0 {
		t.Fatalf("expected score=0, got %d", s)
	}
}

func TestScoreAboveHundred(t *testing.T) {
	// no results -> should be clamped to <= 100
	results := []rules.Result{}
	s := score.Score(results, nil)

	if s != 100 {
		t.Fatalf("expected 100, got %d", s)
	}
}
