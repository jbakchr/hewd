package cmd_test

import (
	"strings"
	"testing"

	"github.com/jbakchr/hewd/internal/rules"
	"github.com/jbakchr/hewd/internal/score"
	"github.com/jbakchr/hewd/pkg/cmd"
)

func TestRenderMarkdown(t *testing.T) {
	out := cmd.DoctorOutput{
		Score: 78,
		Category: score.CategoryScores{
			Documentation: 73,
			Config:        92,
			Structure:     68,
			Overall:       78,
		},
		Results: []score.ScoredRule{
			{
				Result: rules.Result{
					ID:      "DOC_TEST",
					Level:   rules.Info,
					Message: "Test message",
				},
				Category: "documentation",
			},
		},
	}

	md := cmd.RenderMarkdownForTest(out)

	if !strings.Contains(md, "# hewd Report") {
		t.Errorf("missing header")
	}

	if !strings.Contains(md, "Documentation: 73") {
		t.Errorf("missing documentation score")
	}

	if !strings.Contains(md, "DOC_TEST") {
		t.Errorf("missing rule output")
	}
}
