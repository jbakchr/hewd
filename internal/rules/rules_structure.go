package rules

import (
    "strings"

    "github.com/jbakchr/hewd/internal/scan"
)

func init() {
    RegisterRule(RuleLargeRepoNoDocs)
}

func RuleLargeRepoNoDocs(s *scan.Summary) []Result {

    if s.Files < 50 {
        return nil
    }

    // Heuristic: large repo, but no docs folder?
    noDocsDir := true
    for docType, files := range s.DocsFound {
        for _, f := range files {
            if strings.Contains(f, "docs/") {
                noDocsDir = false
            }
        }
        _ = docType // ignore
    }

    if noDocsDir {
        return []Result{{
            ID:      "STRUCT_NO_DOCS_DIR",
            Level:   Info,
            Message: "Repository appears large but has no docs/ directory.",
        }}
    }

    return nil
}