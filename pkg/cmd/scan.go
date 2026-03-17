package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"encoding/json"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newScanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Scan the current directory and summarize the project",
		Long: `Scan the current working directory and produce a project summary,
including detected languages, documentation files, file counts, and structure info.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not get working directory: %w", err)
			}

			summary, err := scanDirectory(cwd)
			if err != nil {
				return err
			}

			// Read flags
			jsonOut, _ := cmd.Flags().GetBool("json")
			yamlOut, _ := cmd.Flags().GetBool("yaml")
			pretty, _ := cmd.Flags().GetBool("pretty")

			// Prevent incompatible flag combinations
			if jsonOut && yamlOut {
				return fmt.Errorf("cannot use --json and --yaml together")
			}

			// JSON output
			if jsonOut {
				return printJSON(summary, pretty)
			}

			// YAML output
			if yamlOut {
				return printYAML(summary)
			}

			// Default human-readable output
			printScanSummary(summary)
			return nil
		},
	}

	cmd.Flags().Bool("json", false, "Output scan result in JSON format")
	cmd.Flags().Bool("yaml", false, "Output scan result in YAML format")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON output")

	return cmd
}

// ScanSummary holds all aggregated data from the scan.
type ScanSummary struct {
	Files         int
	Directories   int
	Languages     map[string]int
	Documentation map[string]bool
}

// scanDirectory recursively walks a directory and collects high-level metadata.
func scanDirectory(root string) (*ScanSummary, error) {
	summary := &ScanSummary{
		Languages: make(map[string]int),
		Documentation: map[string]bool{
			"README.md":       false,
			"LICENSE":         false,
			"CONTRIBUTING.md": false,
		},
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the root directory
		if path == root {
			return nil
		}

		// Count directories and files
		if info.IsDir() {
			summary.Directories++
			return nil
		}

		summary.Files++

		// Detect documentation files
		name := info.Name()
		if _, ok := summary.Documentation[name]; ok {
			summary.Documentation[name] = true
		}

		// Detect languages by extension
		ext := strings.ToLower(filepath.Ext(name))
		if ext != "" {
			ext = ext[1:] // remove leading dot
			summary.Languages[ext]++
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error scanning directory: %w", err)
	}

	return summary, nil
}

// printScanSummary outputs the results in a readable format.
func printScanSummary(s *ScanSummary) {
	fmt.Println("Project Summary:")
	fmt.Printf("  Files:       %d\n", s.Files)
	fmt.Printf("  Directories: %d\n", s.Directories)

	fmt.Println()
	fmt.Println("Languages detected:")

	if len(s.Languages) == 0 {
		fmt.Println("  (none detected)")
	} else {
		// Sort languages alphabetically for stable output
		langs := make([]string, 0, len(s.Languages))
		for lang := range s.Languages {
			langs = append(langs, lang)
		}
		sort.Strings(langs)

		for _, lang := range langs {
			fmt.Printf("  %s (%d files)\n", lang, s.Languages[lang])
		}
	}

	fmt.Println()
	fmt.Println("Documentation:")

	for doc, exists := range s.Documentation {
		if exists {
			fmt.Printf("  %s: present\n", doc)
		} else {
			fmt.Printf("  %s: missing\n", doc)
		}
	}

	fmt.Println()
	fmt.Println("Scan complete.")
}

// printJSON marshals the summary to JSON.
func printJSON(s *ScanSummary, pretty bool) error {
	var data []byte
	var err error

	if pretty {
		data, err = json.MarshalIndent(s, "", "  ")
	} else {
		data, err = json.Marshal(s)
	}

	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	fmt.Println(string(data))
	return nil
}

// printYAML marshals the summary to YAML.
func printYAML(s *ScanSummary) error {
	data, err := yaml.Marshal(s)
	if err != nil {
		return fmt.Errorf("failed to encode YAML: %w", err)
	}

	fmt.Println(string(data))
	return nil
}
