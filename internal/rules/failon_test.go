package rules_test

import (
	"testing"

	"github.com/jbakchr/hewd/internal/rules"
)

func TestSeverityRank(t *testing.T) {
	if rules.SeverityRank(rules.Info) != 1 {
		t.Fatal("info rank expected 1")
	}
	if rules.SeverityRank(rules.Warn) != 2 {
		t.Fatal("warn rank expected 2")
	}
	if rules.SeverityRank(rules.Error) != 3 {
		t.Fatal("error rank expected 3")
	}
}
