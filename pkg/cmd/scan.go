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

//
// ───────────────────────────────────────────────────────────────
//   Language Detection (unchanged, expanded if you want later)
// ───────────────────────────────────────────────────────────────
//

var realLanguages = map[string]string{
	"go":    "Go",
	"js":    "JavaScript",
	"ts":    "TypeScript",
	"py":    "Python",
	"rb":    "Ruby",
	"java":  "Java",
	"rs":    "Rust",
	"c":     "C",
	"h":     "C Header",
	"cpp":   "C++",
	"hpp":   "C++ Header",
	"cs":    "C#",
	"php":   "PHP",
	"swift": "Swift",
	"kt":    "Kotlin",
	"m":     "Objective‑C",
	"mm":    "Objective‑C++",

	// Scripting & Config formats
	"sh":   "Shell",
	"bash": "Bash",
	"zsh":  "Zsh",
	"ps1":  "PowerShell",

	// Markup / Data formats
	"md":       "Markdown",
	"markdown": "Markdown",
	"txt":      "Plain Text",
	"yaml":     "YAML",
	"yml":      "YAML",
	"json":     "JSON",
	"toml":     "TOML",
	"xml":      "XML",
	"html":     "HTML",
}

//
// ───────────────────────────────────────────────────────────────
//   NEW: Documentation and Configuration Asset Detection
// ───────────────────────────────────────────────────────────────
//

// Map of well-known documentation files
var documentationAssets = map[string]string{
	"README.md":          "Project Overview",
	"CONTRIBUTING.md":    "Contribution Guide",
	"CHANGELOG.md":       "Changelog",
	"LICENSE":            "License File",
	"SECURITY.md":        "Security Policy",
	"CODE_OF_CONDUCT.md": "Code of Conduct",
}

// Map of config assets (file → label)
var configAssets = map[string]string{
	"go.mod":             "Go Module",
	"package.json":       "Node Package Manifest",
	"pyproject.toml":     "Python Project Config",
	"Dockerfile":         "Docker Build Config",
	"docker-compose.yml": "Docker Compose Config",
	"openapi.yaml":       "OpenAPI Specification",
	"openapi.yml":        "OpenAPI Specification",
}

//
// ───────────────────────────────────────────────────────────────
//   Cobra command setup
// ───────────────────────────────────────────────────────────────
//

func newScanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Scan the current directory and summarize the project",
		Long: `Scan the current working directory and produce a project summary,
including detected languages, documentation files, file counts, 
configuration files, and structure info.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not get working directory: %w", err)
			}

			summary, err := scanDirectory(cwd)
			if err != nil {
				return err
			}

			jsonOut, _ := cmd.Flags().GetBool("json")
			yamlOut, _ := cmd.Flags().GetBool("yaml")
			pretty, _ := cmd.Flags().GetBool("pretty")

			if jsonOut && yamlOut {
				return fmt.Errorf("cannot use --json and --yaml together")
			}

			if jsonOut {
				return printJSON(summary, pretty)
			}
			if yamlOut {
				return printYAML(summary)
			}

			printScanSummary(summary)
			return nil
		},
	}

	cmd.Flags().Bool("json", false, "Output scan result in JSON format")
	cmd.Flags().Bool("yaml", false, "Output scan result in YAML format")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON output")

	return cmd
}

//
// ───────────────────────────────────────────────────────────────
//   Data structure for scan results
// ───────────────────────────────────────────────────────────────
//

type ScanSummary struct {
	Files         int
	Directories   int
	Languages     map[string]int      `json:"languages" yaml:"languages"`
	Documentation map[string]bool     `json:"documentation" yaml:"documentation"`
	DocsFound     map[string][]string `json:"documentation_files" yaml:"documentation_files"`
	ConfigFiles   map[string][]string `json:"config_files" yaml:"config_files"`
}

//
// ───────────────────────────────────────────────────────────────
//   Scanner Implementation
// ───────────────────────────────────────────────────────────────
//

func scanDirectory(root string) (*ScanSummary, error) {
	summary := &ScanSummary{
		Languages:     make(map[string]int),
		Documentation: make(map[string]bool),
		DocsFound:     make(map[string][]string),
		ConfigFiles:   make(map[string][]string),
	}

	// Initialize doc presence map
	for name := range documentationAssets {
		summary.Documentation[name] = false
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == root {
			return nil
		}

		if info.IsDir() {
			summary.Directories++
			return nil
		}

		summary.Files++
		name := info.Name()

		//
		// Documentation detection
		//
		if docType, ok := documentationAssets[name]; ok {
			summary.Documentation[name] = true
			summary.DocsFound[docType] = append(summary.DocsFound[docType], path)
		}

		//
		// Config detection
		//
		if cfgType, ok := configAssets[name]; ok {
			summary.ConfigFiles[cfgType] = append(summary.ConfigFiles[cfgType], path)
		}

		//
		// Language detection (existing)
		//
		ext := strings.ToLower(filepath.Ext(name))
		if ext != "" {
			ext = ext[1:]
			if langName, ok := realLanguages[ext]; ok {
				summary.Languages[langName]++
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error scanning directory: %w", err)
	}

	return summary, nil
}

//
// ───────────────────────────────────────────────────────────────
//   Output: Pretty Print
// ───────────────────────────────────────────────────────────────
//

func printScanSummary(s *ScanSummary) {
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

//
// ───────────────────────────────────────────────────────────────
//   Output: JSON / YAML
// ───────────────────────────────────────────────────────────────
//

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

func printYAML(s *ScanSummary) error {
	data, err := yaml.Marshal(s)
	if err != nil {
		return fmt.Errorf("failed to encode YAML: %w", err)
	}

	fmt.Println(string(data))
	return nil
}
