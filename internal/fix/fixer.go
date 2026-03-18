package fix

import (
	"os"
	"path/filepath"

	"github.com/jbakchr/hewd/internal/rules"
)

type Fix struct {
	RuleID   string
	Message  string
	FilePath string
}

func DetectFixes(results []rules.Result, root string) []Fix {
	var fixes []Fix

	for _, r := range results {
		switch r.ID {

		case "DOC_CONTRIBUTING_MISSING":
			fixes = append(fixes, Fix{
				RuleID:   r.ID,
				Message:  "Create CONTRIBUTING.md",
				FilePath: filepath.Join(root, "CONTRIBUTING.md"),
			})

		case "DOC_LICENSE_MISSING":
			fixes = append(fixes, Fix{
				RuleID:   r.ID,
				Message:  "Create LICENSE file",
				FilePath: filepath.Join(root, "LICENSE"),
			})

		case "DOC_CHANGELOG_EMPTY":
			fixes = append(fixes, Fix{
				RuleID:   r.ID,
				Message:  "Create CHANGELOG.md skeleton",
				FilePath: filepath.Join(root, "CHANGELOG.md"),
			})

		case "CFG_GITHUB_NO_WORKFLOWS":
			fixes = append(fixes, Fix{
				RuleID:   r.ID,
				Message:  "Create .github/workflows/ci.yml",
				FilePath: filepath.Join(root, ".github/workflows/ci.yml"),
			})

		case "STRUCT_NO_DOCS_DIR_FOR_LARGE_REPO":
			fixes = append(fixes, Fix{
				RuleID:   r.ID,
				Message:  "Create docs/ directory",
				FilePath: filepath.Join(root, "docs"),
			})

		}
	}

	return fixes
}

func ApplyFix(f Fix) error {
	// Directory case
	if filepath.Ext(f.FilePath) == "" {
		return os.MkdirAll(f.FilePath, 0755)
	}

	// File case: do not overwrite existing files
	if _, err := os.Stat(f.FilePath); err == nil {
		return nil
	}

	// Ensure parent directory exists
	_ = os.MkdirAll(filepath.Dir(f.FilePath), 0755)

	content := defaultContentForFix(f.RuleID)
	return os.WriteFile(f.FilePath, []byte(content), 0644)
}

func defaultContentForFix(ruleID string) string {
	switch ruleID {
	case "DOC_CONTRIBUTING_MISSING":
		return "# Contributing\n\nThank you for your interest in contributing!"
	case "DOC_LICENSE_MISSING":
		return "MIT License\n\n<your license here>"
	case "DOC_CHANGELOG_EMPTY":
		return "# Changelog\n\nAll notable changes will be documented here."
	case "CFG_GITHUB_NO_WORKFLOWS":
		return "name: CI\n\non: [push]\n\njobs:\n  build:\n    runs-on: ubuntu-latest\n    steps:\n      - uses: actions/checkout@v4\n      - run: echo \"CI running\"\n"
	case "STRUCT_NO_DOCS_DIR_FOR_LARGE_REPO":
		return ""
	}
	return ""
}
