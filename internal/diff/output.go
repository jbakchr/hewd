package diff

import (
	"time"

	"github.com/jbakchr/hewd/internal/api"
	"github.com/jbakchr/hewd/internal/score"
)

type DiffOutput struct {
	SchemaVersion  int                `json:"schemaVersion" yaml:"schemaVersion"`
	HewdVersion    string             `json:"hewdVersion" yaml:"hewdVersion"`
	GeneratedAt    time.Time          `json:"generatedAt" yaml:"generatedAt"`
	ScoreDelta     int                `json:"scoreDelta" yaml:"scoreDelta"`
	CategoryDeltas map[string]int     `json:"categoryDeltas" yaml:"categoryDeltas"`
	NewIssues      []score.ScoredRule `json:"newIssues" yaml:"newIssues"`
	ResolvedIssues []score.ScoredRule `json:"resolvedIssues" yaml:"resolvedIssues"`
}

func MakeDiffOutput(dr DiffResult, old, new *api.MachineOutput) DiffOutput {
	return DiffOutput{
		SchemaVersion:  old.SchemaVersion, // confirmed equal earlier
		HewdVersion:    new.HewdVersion,
		GeneratedAt:    time.Now(),
		ScoreDelta:     dr.ScoreDelta,
		CategoryDeltas: dr.CategoryDeltas,
		NewIssues:      dr.NewIssues,
		ResolvedIssues: dr.ResolvedIssues,
	}
}
