package diff

import (
	"github.com/jbakchr/hewd/internal/api"
	"github.com/jbakchr/hewd/internal/score"
)

type DiffResult struct {
	ScoreDelta     int
	CategoryDeltas map[string]int
	NewIssues      []score.ScoredRule
	ResolvedIssues []score.ScoredRule
}

func ComputeDiff(old, new *api.MachineOutput) DiffResult {
	dr := DiffResult{
		ScoreDelta:     new.Score - old.Score,
		CategoryDeltas: make(map[string]int),
	}

	// --------------------------------------------------------------------
	// Category deltas (explicit, matching CategoryScores struct)
	// --------------------------------------------------------------------
	dr.CategoryDeltas["documentation"] =
		new.CategoryScores.Documentation - old.CategoryScores.Documentation

	dr.CategoryDeltas["config"] =
		new.CategoryScores.Config - old.CategoryScores.Config

	dr.CategoryDeltas["structure"] =
		new.CategoryScores.Structure - old.CategoryScores.Structure

	// --------------------------------------------------------------------
	// Build maps for fast lookups
	// --------------------------------------------------------------------
	oldMap := make(map[string]score.ScoredRule)
	for _, r := range old.Results {
		oldMap[r.ID] = r
	}

	newMap := make(map[string]score.ScoredRule)
	for _, r := range new.Results {
		newMap[r.ID] = r
	}

	// --------------------------------------------------------------------
	// New issues
	// --------------------------------------------------------------------
	for id, r := range newMap {
		if _, exists := oldMap[id]; !exists {
			dr.NewIssues = append(dr.NewIssues, r)
		}
	}

	// --------------------------------------------------------------------
	// Resolved issues
	// --------------------------------------------------------------------
	for id, r := range oldMap {
		if _, exists := newMap[id]; !exists {
			dr.ResolvedIssues = append(dr.ResolvedIssues, r)
		}
	}

	// --------------------------------------------------------------------
	// Sort
	// --------------------------------------------------------------------
	dr.NewIssues = SortIssues(dr.NewIssues)
	dr.ResolvedIssues = SortIssues(dr.ResolvedIssues)

	return dr
}
