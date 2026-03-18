package rules_test

import (
	"testing"

	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/scan"
)

func TestRuleReadmeMissing(t *testing.T) {
	summary := &scan.Summary{
		Documentation: map[string]bool{
			"README.md": false,
		},
		Languages:   map[string]int{},
		DocsFound:   map[string][]string{},
		ConfigFiles: map[string][]string{},
	}

	results := rules.RuleReadmeMissing(summary)

	if len(results) == 0 {
		t.Fatal("RuleReadmeMissing should trigger")
	}

	if results[0].ID != "DOC_README_MISSING" {
		t.Fatalf("expected DOC_README_MISSING, got %s", results[0].ID)
	}
}
