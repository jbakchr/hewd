package cmd

import (
	"fmt"
	"strings"

	"github.com/jbakchr/hewd/internal/score"
)

func renderMarkdown(out DoctorOutput) string {
	var b strings.Builder

	// Header
	fmt.Fprintf(&b, "# hewd Report\n\n")

	// Overall score
	fmt.Fprintf(&b, "## Project Health Score\n")
	fmt.Fprintf(&b, "**%d / 100**\n\n", out.Score)

	// Category scores
	fmt.Fprintf(&b, "### Category Scores\n")
	fmt.Fprintf(&b, "- Documentation: %d\n", out.Category.Documentation)
	fmt.Fprintf(&b, "- Config: %d\n", out.Category.Config)
	fmt.Fprintf(&b, "- Structure: %d\n", out.Category.Structure)
	fmt.Fprintf(&b, "- Overall: %d\n\n", out.Category.Overall)

	// Group results by category
	grouped := map[string][]score.ScoredRule{}
	for _, r := range out.Results {
		grouped[r.Category] = append(grouped[r.Category], r)
	}

	// Output results
	fmt.Fprintf(&b, "## Issues by Category\n\n")

	for cat, list := range grouped {
		fmt.Fprintf(&b, "### %s\n\n", strings.Title(cat))
		if len(list) == 0 {
			fmt.Fprintf(&b, "(none)\n\n")
			continue
		}

		for _, r := range list {
			msg := r.Message
			if r.File != "" {
				msg += fmt.Sprintf(" *(File: %s)*", r.File)
			}
			fmt.Fprintf(&b, "- **[%s] %s** — %s\n", r.Level, r.ID, msg)
		}
		fmt.Fprintf(&b, "\n")
	}

	return b.String()
}

// Export for tests only
func RenderMarkdownForTest(out DoctorOutput) string { return renderMarkdown(out) }
