package api

import (
	"time"

	"github.com/jbakchr/hewd/internal/score"
)

// MachineOutput defines the canonical, stable, machine-readable output schema
// for hewd. This structure is versioned to enable safe evolution of the format
// without breaking consumers (CI pipelines, dashboards, bots, etc.).
type MachineOutput struct {
	// SchemaVersion increments only when the output format itself changes.
	SchemaVersion int `json:"schemaVersion"`

	// HewdVersion is the version of the hewd CLI, pulled from internal/version.
	HewdVersion string `json:"hewdVersion"`

	// GeneratedAt is the timestamp at which the output was produced.
	GeneratedAt time.Time `json:"generatedAt"`

	// Score is the overall project maturity score (0–100).
	Score int `json:"score"`

	// CategoryScores provides documentation/config/structure scoring breakdown.
	CategoryScores score.CategoryScores `json:"categoryScores"`

	// Results is a full list of rule evaluations, including severity, message,
	// file context, and category (via ScoredRule).
	Results []score.ScoredRule `json:"results"`

	// Fixable enumerates all issues for which hewd can apply an automated fix.
	Fixable []FixableItem `json:"fixable"`
}

// FixableItem represents a fixable hewd rule result, including a recommended
// fix message and the file path that would be created or affected.
type FixableItem struct {
	RuleID   string `json:"ruleId"`
	Message  string `json:"message"`
	FilePath string `json:"filePath"`
}
