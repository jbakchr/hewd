package scan

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jbakchr/hewd/internal/config"
)

// ScanDirectory walks the directory at the given root and produces a Summary.
// It respects include/exclude settings from `.hewd/config.yaml`.
func ScanDirectory(root string) (*Summary, error) {
	// Load config if available
	cfg, _ := config.Load(".") // ignore errors; missing config is allowed

	s := &Summary{
		Languages:     make(map[string]int),
		Documentation: make(map[string]bool),
		DocsFound:     make(map[string][]string),
		ConfigFiles:   make(map[string][]string),
	}

	// Initialize documentation presence map
	for name := range DocumentationAssets {
		s.Documentation[name] = false
	}

	// Preprocess include/exclude patterns
	includes := cfg.Scan.Include
	excludes := cfg.Scan.Exclude

	// Helper: check if any exclude prefix matches this path
	isExcluded := func(rel string) bool {
		for _, ex := range excludes {
			if strings.HasPrefix(rel, ex) {
				return true
			}
		}
		return false
	}

	// Helper: if includes list is not empty, only allow matches
	isIncluded := func(rel string) bool {
		if len(includes) == 0 {
			return true // no includes → allow all
		}
		for _, inc := range includes {
			if strings.HasPrefix(rel, inc) {
				return true
			}
		}
		return false
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		// Compute relative path for include/exclude logic
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		// Never exclude the root folder itself
		if rel == "." {
			return nil
		}

		// Exclusion takes precedence
		if isExcluded(rel) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Inclusion check (only if includes exist)
		if !isIncluded(rel) {
			return nil
		}

		// Count directories
		if info.IsDir() {
			s.Directories++
			return nil
		}

		// Count files
		s.Files++

		name := info.Name()

		// ----------------------------------------------------------------------
		// Documentation detection
		// ----------------------------------------------------------------------
		if docType, ok := DocumentationAssets[name]; ok {
			s.Documentation[name] = true
			s.DocsFound[docType] = append(s.DocsFound[docType], path)
		}

		// ----------------------------------------------------------------------
		// Config detection
		// ----------------------------------------------------------------------
		if cfgType, ok := ConfigAssets[name]; ok {
			s.ConfigFiles[cfgType] = append(s.ConfigFiles[cfgType], path)
		}

		// ----------------------------------------------------------------------
		// Language detection
		// ----------------------------------------------------------------------
		ext := strings.ToLower(filepath.Ext(name))
		if ext != "" {
			ext = ext[1:] // strip leading dot
			if langName, ok := RealLanguages[ext]; ok {
				s.Languages[langName]++
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error scanning directory: %w", err)
	}

	return s, nil
}
