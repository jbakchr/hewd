package cmd

import (
	"fmt"
	"sort"

	"github.com/jbakchr/hewd/internal/cliutils"
	"github.com/jbakchr/hewd/internal/scan"
)

func printScanSummary(sum *scan.Summary) {

	// ===== HEADER =====
	fmt.Printf("%s===== SCAN SUMMARY =====%s\n", cliutils.CyanBold, cliutils.Reset)

	// Basic stats
	fmt.Printf("%sFiles:%s        %d\n", cliutils.WhiteBold, cliutils.Reset, sum.Files)
	fmt.Printf("%sDirectories:%s %d\n\n", cliutils.WhiteBold, cliutils.Reset, sum.Directories)

	// ===== LANGUAGES =====
	fmt.Printf("%s===== LANGUAGES =====%s\n", cliutils.CyanBold, cliutils.Reset)
	if len(sum.Languages) == 0 {
		fmt.Println("  (none)")
	} else {
		langs := make([]string, 0, len(sum.Languages))
		for lang := range sum.Languages {
			langs = append(langs, lang)
		}
		sort.Strings(langs)
		for _, lang := range langs {
			fmt.Printf("  - %s (%d files)\n", lang, sum.Languages[lang])
		}
		fmt.Println()
	}

	// ===== DOCUMENTATION =====
	fmt.Printf("%s===== DOCUMENTATION =====%s\n", cliutils.CyanBold, cliutils.Reset)
	if len(sum.Documentation) == 0 {
		fmt.Println("  (none)")
	} else {
		// Sorted by filename
		files := make([]string, 0, len(sum.Documentation))
		for doc := range sum.Documentation {
			files = append(files, doc)
		}
		sort.Strings(files)
		for _, file := range files {
			state := "missing"
			if sum.Documentation[file] {
				state = "present"
			}
			fmt.Printf("  - %s: %s\n", file, state)
		}
		fmt.Println()
	}

	// ===== DOCUMENTATION FILES FOUND =====
	fmt.Printf("%s===== DOCUMENTATION FILES =====%s\n", cliutils.CyanBold, cliutils.Reset)
	if len(sum.DocsFound) == 0 {
		fmt.Println("  (none)")
	} else {
		types := make([]string, 0, len(sum.DocsFound))
		for t := range sum.DocsFound {
			types = append(types, t)
		}
		sort.Strings(types)

		for _, t := range types {
			fmt.Printf("  %s:\n", t)
			for _, path := range sum.DocsFound[t] {
				fmt.Printf("    - %s\n", path)
			}
		}
		fmt.Println()
	}

	// ===== CONFIG FILES =====
	fmt.Printf("%s===== CONFIG FILES =====%s\n", cliutils.CyanBold, cliutils.Reset)
	if len(sum.ConfigFiles) == 0 {
		fmt.Println("  (none)")
	} else {
		types := make([]string, 0, len(sum.ConfigFiles))
		for t := range sum.ConfigFiles {
			types = append(types, t)
		}
		sort.Strings(types)

		for _, t := range types {
			fmt.Printf("  %s:\n", t)
			for _, path := range sum.ConfigFiles[t] {
				fmt.Printf("    - %s\n", path)
			}
		}
		fmt.Println()
	}
}
