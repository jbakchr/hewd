package api_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/jbakchr/hewd/internal/api"
	"github.com/jbakchr/hewd/internal/score"
)

// Ensure the MachineOutput struct can be created and marshaled.
func TestMachineOutput_Marshal(t *testing.T) {
	m := api.MachineOutput{
		SchemaVersion: 1,
		HewdVersion:   "0.1.0-test",
		GeneratedAt:   time.Now(),
		Score:         88,
		CategoryScores: score.CategoryScores{
			Documentation: 90,
			Config:        80,
			Structure:     95,
			Overall:       88,
		},
		Results: []score.ScoredRule{
			{
				Result:   score.ScoredRule{}.Result, // zero-value OK for marshalling test
				Category: "documentation",
			},
		},
		Fixable: []api.FixableItem{
			{
				RuleID:   "DOC_LICENSE_MISSING",
				Message:  "Create LICENSE",
				FilePath: "/tmp/LICENSE",
			},
		},
	}

	// Test JSON marshalling works without error
	_, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("failed to marshal MachineOutput: %v", err)
	}
}

// Ensure important fields survive JSON round-trip.
func TestMachineOutput_JSONRoundTrip(t *testing.T) {
	m1 := api.MachineOutput{
		SchemaVersion: 1,
		HewdVersion:   "test",
		GeneratedAt:   time.Time{}, // zero value still fine
		Score:         42,
		CategoryScores: score.CategoryScores{
			Documentation: 10,
			Config:        20,
			Structure:     30,
			Overall:       20,
		},
		Fixable: []api.FixableItem{
			{
				RuleID:   "TEST_RULE",
				Message:  "Test fix message",
				FilePath: "test/file",
			},
		},
	}

	// Marshal
	data, err := json.Marshal(m1)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	// Unmarshal
	var m2 api.MachineOutput
	if err := json.Unmarshal(data, &m2); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	// Validate selected fields
	if m2.SchemaVersion != m1.SchemaVersion {
		t.Errorf("SchemaVersion mismatch: got %d, want %d", m2.SchemaVersion, m1.SchemaVersion)
	}
	if m2.Score != m1.Score {
		t.Errorf("Score mismatch: got %d, want %d", m2.Score, m1.Score)
	}
	if len(m2.Fixable) != 1 {
		t.Fatalf("expected 1 fixable item, got %d", len(m2.Fixable))
	}
	if m2.Fixable[0].RuleID != "TEST_RULE" {
		t.Errorf("Fixable.RuleID mismatch: got %s", m2.Fixable[0].RuleID)
	}
}
