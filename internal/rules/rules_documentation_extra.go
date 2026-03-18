package rules

import (
	"os"
	"strings"

	"github.com/jbakchr/hewd/internal/scan"
)

// This file contains additional (non-core) documentation rules.
//
// These rules focus on project documentation maturity, structure,
// and content quality. They complement the core rules defined
// in rules_documentation.go.

func init() {
	RegisterRule("DOCS_EMPTY_FOLDER", RuleDocsFolderExistsButEmpty)
	RegisterRule("DOC_README_NO_USAGE", RuleReadmeMissingUsageSection)
	RegisterRule("DOC_CHANGELOG_EMPTY", RuleChangelogExistsButEmpty)
	RegisterRule("DOC_MANY_MD_NO_DOCS_DIR", RuleTooManyMarkdownFilesWithoutDocsFolder)
}

// -----------------------------------------------------------------------------
// 1. docs/ exists but appears empty
// -----------------------------------------------------------------------------

func RuleDocsFolderExistsButEmpty(s interface{}) []Result {
	summary := s.(*scan.Summary)

	hasDocsDir := false
	hasFilesInside := false

	for _, files := range summary.DocsFound {
		for _, p := range files {
			if strings.Contains(p, "docs/") {
				hasDocsDir = true
				hasFilesInside = true
			}
		}
	}

	// If docs/ exists but no files were actually detected inside it
	if hasDocsDir && !hasFilesInside {
		return []Result{{
			ID:      "DOCS_EMPTY_FOLDER",
			Level:   Warn,
			Message: "A docs/ folder exists but appears empty.",
		}}
	}

	return nil
}

// -----------------------------------------------------------------------------
// 2. README.md missing a “Usage” section
// -----------------------------------------------------------------------------

func RuleReadmeMissingUsageSection(s interface{}) []Result {
	summary := s.(*scan.Summary)

	if !summary.Documentation["README.md"] {
		return nil // Another rule already checks for missing README.
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
	if !strings.Contains(content, "usage") {
		return []Result{{
			ID:      "DOC_README_NO_USAGE",
			Level:   Info,
			Message: "README.md does not contain a 'Usage' section.",
			File:    paths[0],
		}}
	}

	return nil
}

// -----------------------------------------------------------------------------
// 3. CHANGELOG.md exists but is empty or extremely short
// -----------------------------------------------------------------------------

func RuleChangelogExistsButEmpty(s interface{}) []Result {
	summary := s.(*scan.Summary)

	if !summary.Documentation["CHANGELOG.md"] {
		return nil
	}

	paths := summary.DocsFound["Changelog"]
	if len(paths) == 0 {
		return nil
	}

	data, err := os.ReadFile(paths[0])
	if err != nil {
		return nil
	}

	if len(data) < 20 {
		return []Result{{
			ID:      "DOC_CHANGELOG_EMPTY",
			Level:   Warn,
			Message: "CHANGELOG.md exists but appears empty or too short.",
			File:    paths[0],
		}}
	}

	return nil
}

// -----------------------------------------------------------------------------
// 4. Many markdown files but no docs/ folder exists
// -----------------------------------------------------------------------------

func RuleTooManyMarkdownFilesWithoutDocsFolder(s interface{}) []Result {
	summary := s.(*scan.Summary)

	mdCount := 0
	hasDocsDir := false

	// Detect docs/ folder
	for _, list := range summary.DocsFound {
		for _, f := range list {
			if strings.Contains(f, "docs/") {
				hasDocsDir = true
			}
		}
	}

	// Count Markdown files using language detection
	if count, ok := summary.Languages["Markdown"]; ok {
		mdCount = count
	}

	if mdCount > 10 && !hasDocsDir {
		return []Result{{
			ID:      "DOC_MANY_MD_NO_DOCS_DIR",
			Level:   Info,
			Message: "Repository contains many markdown files but no docs/ directory.",
		}}
	}

	return nil
}
