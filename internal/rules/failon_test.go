package rules_test

import (
	"testing"

	"github.com/jbakchr/hewd/internal/rules"
)

func TestSeverityRank(t *testing.T) {
	if rules.SeverityRank(rules.Info) != 1 {
		t.Fatal("info should be rank 1")
	}
	if rules.SeverityRank(rules.Warn) != 2 {
		t.Fatal("warn should be rank 2")
	}
	if rules.SeverityRank(rules.Error) != 3 {
		t.Fatal("error should be rank 3")
	}
}
