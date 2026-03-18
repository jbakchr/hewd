package rules

import (
	"os"
	"strings"

	"github.com/jbakchr/hewd/internal/scan"
)

func init() {
	RegisterRule("STRUCT_MULTI_LANG_NO_ARCH", "structure", RuleManyLanguagesNoArchitectureDoc)
	RegisterRule("STRUCT_NO_LICENSE_REFERENCE", "structure", RuleLargeRepoNoLicenseReference)
	RegisterRule("STRUCT_README_TOO_SMALL", "structure", RuleRootReadmeVerySmall)
}

// 1. Multi-language repo but no architecture docs
func RuleManyLanguagesNoArchitectureDoc(s interface{}) []Result {
	sum := s.(*scan.Summary)

	if len(sum.Languages) < 2 {
		return nil
	}

	hasArchitectureDoc := false

	for _, files := range sum.DocsFound {
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

// 2. Large repo but README does not reference LICENSE
func RuleLargeRepoNoLicenseReference(s interface{}) []Result {
	sum := s.(*scan.Summary)

	if !sum.Documentation["README.md"] {
		return nil
	}
	if sum.Files < 50 {
		return nil
	}

	paths := sum.DocsFound["Project Overview"]
	if len(paths) == 0 {
		return nil
	}

	data, err := os.ReadFile(paths[0])
	if err != nil {
		return nil
	}

	content := strings.ToLower(string(data))

	if sum.Documentation["LICENSE"] && !strings.Contains(content, "license") {
		return []Result{{
			ID:      "STRUCT_NO_LICENSE_REFERENCE",
			Level:   Info,
			Message: "Large project: README exists but does not reference LICENSE.",
			File:    paths[0],
		}}
	}

	return nil
}

// 3. README.md appears too small
func RuleRootReadmeVerySmall(s interface{}) []Result {
	sum := s.(*scan.Summary)

	if !sum.Documentation["README.md"] {
		return nil
	}

	paths := sum.DocsFound["Project Overview"]
	if len(paths) == 0 {
		return nil
	}

	data, err := os.ReadFile(paths[0])
	if err != nil {
		return nil
	}

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
