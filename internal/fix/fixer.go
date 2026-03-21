package fix

import (
	"os"
	"path/filepath"

	"github.com/jbakchr/hewd/internal/cliutils"
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
				Message:  "create CONTRIBUTING.md",
				FilePath: filepath.Join(root, "CONTRIBUTING.md"),
			})

		case "DOC_LICENSE_MISSING":
			fixes = append(fixes, Fix{
				RuleID:   r.ID,
				Message:  "create LICENSE file",
				FilePath: filepath.Join(root, "LICENSE"),
			})

		case "DOC_CHANGELOG_EMPTY":
			fixes = append(fixes, Fix{
				RuleID:   r.ID,
				Message:  "create CHANGELOG.md skeleton",
				FilePath: filepath.Join(root, "CHANGELOG.md"),
			})

		case "CFG_GITHUB_NO_WORKFLOWS":
			fixes = append(fixes, Fix{
				RuleID:   r.ID,
				Message:  "create .github/workflows/ci.yml",
				FilePath: filepath.Join(root, ".github/workflows/ci.yml"),
			})

		case "STRUCT_NO_DOCS_DIR_FOR_LARGE_REPO":
			fixes = append(fixes, Fix{
				RuleID:   r.ID,
				Message:  "create docs directory",
				FilePath: filepath.Join(root, "docs"),
			})
		}
	}

	return fixes
}

func ApplyFix(f Fix) error {
	// Directory case (directory has no extension)
	if filepath.Ext(f.FilePath) == "" {
		if err := os.MkdirAll(f.FilePath, 0755); err != nil {
			return cliutils.ErrHint(
				"failed to create directory",
				"check permissions or run with appropriate user rights",
			)
		}
		return nil
	}

	// File case — ensure file doesn’t already exist
	if _, err := os.Stat(f.FilePath); err == nil {
		// File already exists — non-error but helpful information
		return cliutils.ErrHint(
			"cannot apply fix because file already exists",
			"remove or rename the existing file before applying this fix",
		)
	}

	// Ensure parent directory exists
	if err := os.MkdirAll(filepath.Dir(f.FilePath), 0755); err != nil {
		return cliutils.ErrHint(
			"failed to prepare parent directory",
			"check directory permissions or adjust your repo layout",
		)
	}

	content := defaultContentForFix(f.RuleID)
	if err := os.WriteFile(f.FilePath, []byte(content), 0644); err != nil {
		return cliutils.ErrHint(
			"failed to write file",
			"ensure you have write permissions for this path",
		)
	}

	return nil
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
		return `name: CI

on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: echo "CI running"
`

	case "STRUCT_NO_DOCS_DIR_FOR_LARGE_REPO":
		return ""
	}
	return ""
}
