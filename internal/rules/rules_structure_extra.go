package rules

import (
	"os"
	"strings"

	"github.com/jbakchr/hewd/internal/scan"
)

func init() {
	RegisterRule(RuleManyLanguagesNoArchitectureDoc)
	RegisterRule(RuleLargeRepoNoLicenseReference)
	RegisterRule(RuleRootReadmeVerySmall)
}

// 8. Multiple languages but no architecture docs
func RuleManyLanguagesNoArchitectureDoc(s *scan.Summary) []Result {
	if len(s.Languages) < 2 {
		return nil
	}

	// detect docs/architecture or similar
	found := false
	for _, files := range s.DocsFound {
		for _, f := range files {
			lowered := strings.ToLower(f)
			if strings.Contains(lowered, "architecture") ||
				strings.Contains(lowered, "design") {
				found = true
			}
		}
	}

	if !found {
		return []Result{{
			ID:      "STRUCT_MULTI_LANG_NO_ARCH",
			Level:   Info,
			Message: "Multiple languages detected but no architecture/design docs found.",
		}}
	}

	return nil
}

// 9. Large repo but README does not reference LICENSE
func RuleLargeRepoNoLicenseReference(s *scan.Summary) []Result {
	if !s.Documentation["README.md"] {
		return nil
	}

	if s.Files < 50 {
		return nil
	}

	readmePath := s.DocsFound["Project Overview"]
	if len(readmePath) == 0 {
		return nil
	}

	data, err := os.ReadFile(readmePath[0])
	if err != nil {
		return nil
	}

	text := strings.ToLower(string(data))
	if !strings.Contains(text, "license") && s.Documentation["LICENSE"] {
		return []Result{{
			ID:      "STRUCT_README_NO_LICENSE",
			Level:   Info,
			Message: "Large project: README found but does not reference LICENSE.",
			File:    readmePath[0],
		}}
	}

	return nil
}

// 10. README is too small (heuristic)
func RuleRootReadmeVerySmall(s *scan.Summary) []Result {
	if !s.Documentation["README.md"] {
		return nil
	}

	paths := s.DocsFound["Project Overview"]
	if len(paths) == 0 {
		return nil
	}

	data, err := os.ReadFile(paths[0])
	if err != nil {
		return nil
	}

	if len(data) < 80 {
		return []Result{{
			ID:      "DOC_README_TOO_SMALL",
			Level:   Warn,
			Message: "README.md appears very small; consider expanding project details.",
			File:    paths[0],
		}}
	}

	return nil
}
