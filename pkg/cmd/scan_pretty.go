package cmd

import (
	"fmt"

	"github.com/jbakchr/hewd/internal/scan"
)

func printScanSummary(s *scan.Summary) {
	fmt.Println("Project Summary:")
	fmt.Printf("  Files:       %d\n", s.Files)
	fmt.Printf("  Directories: %d\n", s.Directories)

	fmt.Println()
	fmt.Println("Languages detected:")
	if len(s.Languages) == 0 {
		fmt.Println("  (none detected)")
	} else {
		for lang, count := range s.Languages {
			fmt.Printf("  %s (%d files)\n", lang, count)
		}
	}

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
