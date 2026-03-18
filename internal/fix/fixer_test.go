package fix_test

import (
	"testing"

	"github.com/jbakchr/hewd/internal/fix"
	"github.com/jbakchr/hewd/internal/rules"
)

func TestDetectFixes(t *testing.T) {
	results := []rules.Result{
		{ID: "DOC_CONTRIBUTING_MISSING"},
	}

	fixes := fix.DetectFixes(results, "/tmp")

	if len(fixes) != 1 {
		t.Fatal("expected 1 fix")
	}

	if fixes[0].FilePath != "/tmp/CONTRIBUTING.md" {
		t.Fatalf("unexpected file path %s", fixes[0].FilePath)
	}
}
