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
	cfg, _ := config.Load(root) // missing config is not an error

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

	// Normalize include/exclude settings
	includes := normalizePaths(cfg.Scan.Include)
	excludes := normalizePaths(cfg.Scan.Exclude)

	err := filepath.Walk(root, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		// Compute relative path consistently
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)

		// Skip root
		if rel == "." {
			return nil
		}

		// ------------------------------------------------------------------
		// Always exclude the .hewd/ directory (internal metadata)
		// ------------------------------------------------------------------
		if strings.HasPrefix(rel, ".hewd/") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// ------------------------------------------------------------------
		// Exclude directories listed in config
		// ------------------------------------------------------------------
		if shouldExclude(rel, excludes) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// ------------------------------------------------------------------
		// Include logic: if includes list is non-empty, then the file/dir
		// MUST start with one of the include prefixes.
		// ------------------------------------------------------------------
		if len(includes) > 0 && !shouldInclude(rel, includes) {
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
			ext = ext[1:] // strip dot
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

// -----------------------------------------------------------------------------
// Helpers
// -----------------------------------------------------------------------------

// normalizePaths ensures all paths use forward slashes for consistent prefix checks.
func normalizePaths(paths []string) []string {
	var out []string
	for _, p := range paths {
		p = filepath.ToSlash(p)
		p = strings.TrimPrefix(p, "./")
		out = append(out, p)
	}
	return out
}

// shouldExclude checks if a relative path matches any exclude prefix.
func shouldExclude(rel string, excludes []string) bool {
	for _, ex := range excludes {
		if rel == ex || strings.HasPrefix(rel, ex+"/") {
			return true
		}
	}
	return false
}

// shouldInclude checks if a relative path matches at least one include prefix.
func shouldInclude(rel string, includes []string) bool {
	for _, inc := range includes {
		if rel == inc || strings.HasPrefix(rel, inc+"/") {
			return true
		}
	}
	return false
}
