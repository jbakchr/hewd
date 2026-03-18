package badge_test

import (
	"strings"
	"testing"

	"github.com/jbakchr/hewd/internal/badge"
)

func TestBadgeContainsScore(t *testing.T) {
	svg := badge.Generate(85)

	if !strings.Contains(svg, ">85<") {
		t.Fatalf("expected badge to contain score 85")
	}
}

func TestBadgeColorRanges(t *testing.T) {
	if badge.Color(95) != "#4c1" {
		t.Errorf("expected bright green for 95")
	}
	if badge.Color(80) != "#97CA00" {
		t.Errorf("expected green for 80")
	}
	if badge.Color(30) != "#e05d44" {
		t.Errorf("expected red for 30")
	}
}
