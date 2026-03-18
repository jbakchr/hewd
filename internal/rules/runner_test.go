package rules_test

import (
	"testing"

	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/scan"
)

// A fake summary for simple rule testing
func fakeSummary(missingReadme bool) *scan.Summary {
	return &scan.Summary{
		Documentation: map[string]bool{
			"README.md": missingReadme == false,
			"LICENSE":   true,
		},
		Languages:   map[string]int{},
		DocsFound:   map[string][]string{},
		ConfigFiles: map[string][]string{},
	}
}

func TestRuleDisabling(t *testing.T) {
	cfg := &config.Config{
		Rules: map[string]bool{
			"DOC_README_MISSING": false, // disable rule
		},
	}

	s := fakeSummary(true)
	results := rules.RunAll(s, cfg)

	for _, r := range results {
		if r.ID == "DOC_README_MISSING" {
			t.Fatalf("rule should be disabled")
		}
	}
}

func TestSeverityOverride(t *testing.T) {
	cfg := &config.Config{
		Weights: map[string]int{
			"DOC_LICENSE_MISSING": 3, // override to error
		},
	}

	s := &scan.Summary{
		Documentation: map[string]bool{
			"README.md": true,
			"LICENSE":   false,
		},
		Languages:   map[string]int{},
		DocsFound:   map[string][]string{},
		ConfigFiles: map[string][]string{},
	}

	results := rules.RunAll(s, cfg)

	found := false
	for _, r := range results {
		if r.ID == "DOC_LICENSE_MISSING" {
			found = true
			if r.Level != rules.Error {
				t.Errorf("expected overridden severity to be error, got %s", r.Level)
			}
		}
	}

	if !found {
		t.Fatalf("expected DOC_LICENSE_MISSING result")
	}
}
