package rules

import (
	"os"
	"strings"

	"github.com/jbakchr/hewd/internal/scan"
)

func init() {
	RegisterRule(RuleDocsFolderExistsButEmpty)
	RegisterRule(RuleReadmeMissingUsageSection)
	RegisterRule(RuleChangelogExistsButEmpty)
	RegisterRule(RuleTooManyMarkdownFilesWithoutDocsFolder)
}

// 1. docs/ exists but empty
func RuleDocsFolderExistsButEmpty(s *scan.Summary) []Result {
	hasDocsDir := false
	hasFilesInside := false

	for _, files := range s.DocsFound {
		for _, p := range files {
			if strings.Contains(p, "docs/") {
				hasDocsDir = true
				hasFilesInside = true
			}
		}
	}

	// If directory exists but nothing inside is detected
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
func RuleReadmeMissingUsageSection(s *scan.Summary) []Result {
	if !s.Documentation["README.md"] {
		return nil // handled by other rule
	}

	// Check if README file path is known
	paths := s.DocsFound["Project Overview"]
	if len(paths) == 0 {
		return nil
	}

	// Load content
	data, err := os.ReadFile(paths[0])
	if err != nil {
		return nil // reading errors are not the concern of this rule
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

// 3. CHANGELOG exists but empty / very small
func RuleChangelogExistsButEmpty(s *scan.Summary) []Result {
	if !s.Documentation["CHANGELOG.md"] {
		return nil
	}

	paths := s.DocsFound["Changelog"]
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
			Message: "CHANGELOG exists but appears empty or too short.",
			File:    paths[0],
		}}
	}

	return nil
}

// 4. Many markdown files but no docs folder
func RuleTooManyMarkdownFilesWithoutDocsFolder(s *scan.Summary) []Result {
	mdCount := 0
	hasDocsFolder := false

	for _, list := range s.DocsFound {
		for _, f := range list {
			if strings.Contains(f, "docs/") {
				hasDocsFolder = true
			}
		}
	}

	// Count Markdown via language detection
	if count, ok := s.Languages["Markdown"]; ok {
		mdCount = count
	}

	if mdCount > 10 && !hasDocsFolder {
		return []Result{{
			ID:      "DOC_MANY_MD_NO_DOCS_DIR",
			Level:   Info,
			Message: "Repository contains many markdown files but no docs/ folder.",
		}}
	}

	return nil
}
