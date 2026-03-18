package rules

import (
	"os"
	"strings"

	"github.com/jbakchr/hewd/internal/scan"
)

func init() {
	RegisterRule("DOCS_EMPTY_FOLDER", "documentation", RuleDocsFolderExistsButEmpty)
	RegisterRule("DOC_README_NO_USAGE", "documentation", RuleReadmeMissingUsageSection)
	RegisterRule("DOC_CHANGELOG_EMPTY", "documentation", RuleChangelogExistsButEmpty)
	RegisterRule("DOC_MANY_MD_NO_DOCS_DIR", "documentation", RuleTooManyMarkdownFilesWithoutDocsFolder)
}

// 1. docs/ exists but appears empty
func RuleDocsFolderExistsButEmpty(s interface{}) []Result {
	sum := s.(*scan.Summary)

	hasDocsDir := false
	hasFilesInside := false

	for _, files := range sum.DocsFound {
		for _, p := range files {
			if strings.Contains(p, "docs/") {
				hasDocsDir = true
				hasFilesInside = true
			}
		}
	}

	if hasDocsDir && !hasFilesInside {
		return []Result{{
			ID:      "DOCS_EMPTY_FOLDER",
			Level:   Warn,
			Message: "A docs/ folder exists but appears empty.",
		}}
	}

	return nil
}

// 2. README.md missing a “Usage” section
func RuleReadmeMissingUsageSection(s interface{}) []Result {
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

// 3. CHANGELOG.md exists but appears empty or tiny
func RuleChangelogExistsButEmpty(s interface{}) []Result {
	sum := s.(*scan.Summary)

	if !sum.Documentation["CHANGELOG.md"] {
		return nil
	}

	paths := sum.DocsFound["Changelog"]
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

// 4. Many Markdown files but no docs/ folder
func RuleTooManyMarkdownFilesWithoutDocsFolder(s interface{}) []Result {
	sum := s.(*scan.Summary)

	mdCount := 0
	hasDocsDir := false

	for _, list := range sum.DocsFound {
		for _, f := range list {
			if strings.Contains(f, "docs/") {
				hasDocsDir = true
			}
		}
	}

	if count, ok := sum.Languages["Markdown"]; ok {
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
