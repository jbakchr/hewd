package rules_test

import (
	"testing"

	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/scan"
)

func fakeSummary(hasReadme bool) *scan.Summary {
	return &scan.Summary{
		Documentation: map[string]bool{
			"README.md": hasReadme,
			"LICENSE":   true,
		},
		Languages:   map[string]int{},
		DocsFound:   map[string][]string{},
		ConfigFiles: map[string][]string{},
	}
}

// --------------------------------------------------
// Rule disabling
// --------------------------------------------------

func TestRuleDisabling(t *testing.T) {
	cfg := &config.Config{
		Rules: map[string]bool{
			"DOC_README_MISSING": false,
		},
	}

	s := fakeSummary(false)

	results := rules.RunAll(s, cfg, nil, nil)

	for _, r := range results {
		if r.ID == "DOC_README_MISSING" {
			t.Fatalf("DOC_README_MISSING should be disabled")
		}
	}
}

// --------------------------------------------------
// Severity override
// --------------------------------------------------

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

	results := rules.RunAll(s, cfg, nil, nil)

	found := false
	for _, r := range results {
		if r.ID == "DOC_LICENSE_MISSING" {
			found = true
			if r.Level != rules.Error {
				t.Fatalf("expected overridden severity to be error, got %s", r.Level)
			}
		}
	}

	if !found {
		t.Fatal("did not find DOC_LICENSE_MISSING result")
	}
}

// --------------------------------------------------
// Category filtering (only)
// --------------------------------------------------

func TestCategoryFiltering_Only(t *testing.T) {
	s := fakeSummary(false)

	only := []string{"documentation"}
	except := []string{}

	results := rules.RunAll(s, nil, only, except)

	for _, r := range results {
		if rules.CategoryForRule(r.ID) != "documentation" {
			t.Fatalf("expected only documentation rules, got %s", r.ID)
		}
	}
}

// --------------------------------------------------
// Category filtering (except)
// --------------------------------------------------

func TestCategoryFiltering_Except(t *testing.T) {
	s := fakeSummary(false)

	only := []string{}
	except := []string{"documentation"}

	results := rules.RunAll(s, nil, only, except)

	for _, r := range results {
		if rules.CategoryForRule(r.ID) == "documentation" {
			t.Fatalf("documentation rule should have been excluded: %s", r.ID)
		}
	}
}
