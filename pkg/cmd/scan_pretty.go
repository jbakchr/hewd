package cmd

import (
	"fmt"
	"sort"

	iscan "github.com/jbakchr/hewd/internal/scan"
)

//
// ───────────────────────────────────────────────────────────────
//   Output: Pretty Print
// ───────────────────────────────────────────────────────────────
//

func printScanSummary(s *iscan.Summary) {
	fmt.Println("Project Summary:")
	fmt.Printf("  Files:       %d\n", s.Files)
	fmt.Printf("  Directories: %d\n", s.Directories)

	//
	// Languages
	//
	fmt.Println()
	fmt.Println("Languages detected:")

	if len(s.Languages) == 0 {
		fmt.Println("  (none detected)")
	} else {
		langs := make([]string, 0, len(s.Languages))
		for lang := range s.Languages {
			langs = append(langs, lang)
		}
		sort.Strings(langs)

		for _, lang := range langs {
			fmt.Printf("  %s (%d files)\n", lang, s.Languages[lang])
		}
	}

	//
	// Documentation
	//
	fmt.Println()
	fmt.Println("Documentation presence:")
	for doc, exists := range s.Documentation {
		if exists {
			fmt.Printf("  %s: present\n", doc)
		} else {
			fmt.Printf("  %s: missing\n", doc)
		}
	}

	fmt.Println()
	fmt.Println("Documentation files (detailed):")
	if len(s.DocsFound) == 0 {
		fmt.Println("  (none found)")
	} else {
		for docType, files := range s.DocsFound {
			fmt.Printf("  %s:\n", docType)
			for _, file := range files {
				fmt.Printf("    - %s\n", file)
			}
		}
	}

	//
	// Configuration
	//
	fmt.Println()
	fmt.Println("Configuration files:")
	if len(s.ConfigFiles) == 0 {
		fmt.Println("  (none found)")
	} else {
		for cfgType, files := range s.ConfigFiles {
			fmt.Printf("  %s:\n", cfgType)
			for _, file := range files {
				fmt.Printf("    - %s\n", file)
			}
		}
	}

	fmt.Println()
	fmt.Println("Scan complete.")
}
