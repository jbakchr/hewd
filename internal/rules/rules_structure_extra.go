package rules

import (
	"os"
	"strings"
	"time"

	"github.com/jbakchr/hewd/internal/scan"
)

func init() {
	RegisterRule("STRUCT_NO_DOCS_DIR_FOR_LARGE_REPO", "structure", RuleLargeRepoNoDocsDir)
	RegisterRule("STRUCT_STALE_CHANGELOG", "structure", RuleStaleChangelog)
	RegisterRule("STRUCT_STALE_README", "structure", RuleStaleReadme)
}

// 1. Large repo but no docs/ directory
func RuleLargeRepoNoDocsDir(s interface{}) []Result {
	sum := s.(*scan.Summary)

	if sum.Files < 50 {
		return nil
	}

	hasDocsDir := false
	for _, files := range sum.DocsFound {
		for _, p := range files {
			if strings.Contains(p, "docs/") {
				hasDocsDir = true
				break
			}
		}
	}

	if !hasDocsDir {
		return []Result{{
			ID:      "STRUCT_NO_DOCS_DIR_FOR_LARGE_REPO",
			Level:   Info,
			Message: "Large repository detected but no docs/ directory found.",
		}}
	}

	return nil
}

// 2. Changelog exists but stale (not updated in a long time)
func RuleStaleChangelog(s interface{}) []Result {
	sum := s.(*scan.Summary)

	if !sum.Documentation["CHANGELOG.md"] {
		return nil
	}

	paths := sum.DocsFound["Changelog"]
	if len(paths) == 0 {
		return nil
	}

	path := paths[0]

	info, err := os.Stat(path)
	if err != nil {
		return nil
	}

	const staleDays = 180
	age := daysSince(info.ModTime())

	if age > staleDays {
		return []Result{{
			ID:      "STRUCT_STALE_CHANGELOG",
			Level:   Warn,
			Message: "CHANGELOG.md appears stale (no updates for a long time).",
			File:    path,
		}}
	}

	return nil
}

// 3. README.md stale
func RuleStaleReadme(s interface{}) []Result {
	sum := s.(*scan.Summary)

	if !sum.Documentation["README.md"] {
		return nil
	}

	paths := sum.DocsFound["Project Overview"]
	if len(paths) == 0 {
		return nil
	}

	path := paths[0]

	info, err := os.Stat(path)
	if err != nil {
		return nil
	}

	const staleDays = 365
	age := daysSince(info.ModTime())

	if age > staleDays {
		return []Result{{
			ID:      "STRUCT_STALE_README",
			Level:   Info,
			Message: "README.md appears stale (no updates for a long time).",
			File:    path,
		}}
	}

	return nil
}

func daysSince(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
