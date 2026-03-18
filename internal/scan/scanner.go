package scan

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jbakchr/hewd/internal/config"
)

// ScanDirectory walks the project directory and builds a Summary.
// It respects include/exclude lists from .hewd/config.yaml.
func ScanDirectory(root string) (*Summary, error) {
	cfg, _ := config.Load(root)

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

	includes := normalizePaths(cfg.Scan.Include)
	excludes := normalizePaths(cfg.Scan.Exclude)

	err := filepath.Walk(root, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)

		if rel == "." {
			return nil
		}

		// Always exclude .hewd
		if strings.HasPrefix(rel, ".hewd/") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Exclusion rule
		if shouldExclude(rel, excludes) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Inclusion rule
		if len(includes) > 0 && !shouldInclude(rel, includes) {
			return nil
		}

		if info.IsDir() {
			s.Directories++
			return nil
		}

		// It's a file — count it
		s.Files++

		name := info.Name()

		// Documentation detection
		if docType, ok := DocumentationAssets[name]; ok {
			s.Documentation[name] = true
			s.DocsFound[docType] = append(s.DocsFound[docType], path)
		}

		// Config detection
		if cfgType, ok := ConfigAssets[name]; ok {
			s.ConfigFiles[cfgType] = append(s.ConfigFiles[cfgType], path)
		}

		// Language detection
		ext := strings.ToLower(filepath.Ext(name))
		if ext != "" {
			ext = ext[1:]
			if lang, ok := RealLanguages[ext]; ok {
				s.Languages[lang]++
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error scanning directory: %w", err)
	}

	return s, nil
}

// -----------------------------------------------------------------------------
// Helper functions
// -----------------------------------------------------------------------------

func normalizePaths(paths []string) []string {
	var out []string
	for _, p := range paths {
		p = filepath.ToSlash(p)
		p = strings.TrimPrefix(p, "./")
		out = append(out, p)
	}
	return out
}

func shouldExclude(rel string, excludes []string) bool {
	for _, ex := range excludes {
		if rel == ex || strings.HasPrefix(rel, ex+"/") {
			return true
		}
	}
	return false
}

func shouldInclude(rel string, includes []string) bool {
	for _, inc := range includes {
		if rel == inc || strings.HasPrefix(rel, inc+"/") {
			return true
		}
	}
	return false
}
