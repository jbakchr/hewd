package rules

import (
	"os"
	"strings"
	"time"

	"github.com/jbakchr/hewd/internal/scan"
)

// This file contains additional (non-core) structural and metadata rules.
// These focus on heuristic project maturity indicators, large-repo hygiene,
// and basic documentation expectations for real-world repositories.

func init() {
	RegisterRule("STRUCT_NO_DOCS_DIR_FOR_LARGE_REPO", RuleLargeRepoNoDocsDir)
	RegisterRule("STRUCT_STALE_CHANGELOG", RuleStaleChangelog)
	RegisterRule("STRUCT_STALE_README", RuleStaleReadme)
}

// -----------------------------------------------------------------------------
// 1. Large repo but no docs/ directory
// -----------------------------------------------------------------------------

func RuleLargeRepoNoDocsDir(s interface{}) []Result {
	summary := s.(*scan.Summary)

	if summary.Files < 50 {
		return nil // heuristic: only care for larger projects
	}

	hasDocsDir := false
	for _, files := range summary.DocsFound {
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

// -----------------------------------------------------------------------------
// 2. Changelog exists but is stale (last modified long ago)
// -----------------------------------------------------------------------------

func RuleStaleChangelog(s interface{}) []Result {
	summary := s.(*scan.Summary)

	if !summary.Documentation["CHANGELOG.md"] {
		return nil
	}

	paths := summary.DocsFound["Changelog"]
	if len(paths) == 0 {
		return nil
	}

	path := paths[0]

	info, err := os.Stat(path)
	if err != nil {
		return nil
	}

	// Heuristic: warn if changelog hasn't been updated for ~180 days.
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

// -----------------------------------------------------------------------------
// 3. README.md appears stale (not updated recently)
// -----------------------------------------------------------------------------

func RuleStaleReadme(s interface{}) []Result {
	summary := s.(*scan.Summary)

	if !summary.Documentation["README.md"] {
		return nil
	}

	paths := summary.DocsFound["Project Overview"]
	if len(paths) == 0 {
		return nil
	}

	path := paths[0]

	info, err := os.Stat(path)
	if err != nil {
		return nil
	}

	// Heuristic: warn if README hasn't been updated for ~365 days.
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

// -----------------------------------------------------------------------------
// Helpers
// -----------------------------------------------------------------------------

func daysSince(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
