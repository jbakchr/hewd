package rules

import (
	"os"
	"strings"

	"github.com/jbakchr/hewd/internal/scan"
)

// Core structural rules for repo maturity, documentation expectations,
// and multi-language project architecture.

func init() {
	RegisterRule("STRUCT_MULTI_LANG_NO_ARCH", RuleManyLanguagesNoArchitectureDoc)
	RegisterRule("STRUCT_NO_LICENSE_REFERENCE", RuleLargeRepoNoLicenseReference)
	RegisterRule("STRUCT_README_TOO_SMALL", RuleRootReadmeVerySmall)
}

// -----------------------------------------------------------------------------
// 1. Multiple languages but no architecture documentation
// -----------------------------------------------------------------------------

func RuleManyLanguagesNoArchitectureDoc(s interface{}) []Result {
	summary := s.(*scan.Summary)

	if len(summary.Languages) < 2 {
		return nil
	}

	// check for docs named 'architecture', 'design', etc.
	hasArchitectureDoc := false

	for _, files := range summary.DocsFound {
		for _, f := range files {
			l := strings.ToLower(f)
			if strings.Contains(l, "architecture") || strings.Contains(l, "design") {
				hasArchitectureDoc = true
				break
			}
		}
	}

	if !hasArchitectureDoc {
		return []Result{{
			ID:      "STRUCT_MULTI_LANG_NO_ARCH",
			Level:   Info,
			Message: "Multiple languages detected but no architecture/design docs found.",
		}}
	}

	return nil
}

// -----------------------------------------------------------------------------
// 2. Large repo but README does not reference LICENSE
// -----------------------------------------------------------------------------

func RuleLargeRepoNoLicenseReference(s interface{}) []Result {
	summary := s.(*scan.Summary)

	if !summary.Documentation["README.md"] {
		return nil // another rule handles missing README
	}
	if summary.Files < 50 {
		return nil // heuristic: only warn for larger projects
	}

	paths := summary.DocsFound["Project Overview"]
	if len(paths) == 0 {
		return nil
	}

	data, err := os.ReadFile(paths[0])
	if err != nil {
		return nil
	}

	content := strings.ToLower(string(data))

	// Only warn if LICENSE exists but README does not reference it
	if summary.Documentation["LICENSE"] && !strings.Contains(content, "license") {
		return []Result{{
			ID:      "STRUCT_NO_LICENSE_REFERENCE",
			Level:   Info,
			Message: "Large project: README exists but does not reference LICENSE.",
			File:    paths[0],
		}}
	}

	return nil
}

// -----------------------------------------------------------------------------
// 3. README.md appears too small (heuristic)
// -----------------------------------------------------------------------------

func RuleRootReadmeVerySmall(s interface{}) []Result {
	summary := s.(*scan.Summary)

	if !summary.Documentation["README.md"] {
		return nil
	}

	paths := summary.DocsFound["Project Overview"]
	if len(paths) == 0 {
		return nil
	}

	data, err := os.ReadFile(paths[0])
	if err != nil {
		return nil
	}

	// Very small README (heuristic threshold)
	if len(data) < 80 {
		return []Result{{
			ID:      "STRUCT_README_TOO_SMALL",
			Level:   Warn,
			Message: "README.md appears very small; consider expanding project details.",
			File:    paths[0],
		}}
	}

	return nil
}
