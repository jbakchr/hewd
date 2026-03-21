package scan

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jbakchr/hewd/internal/cliutils"
	"github.com/jbakchr/hewd/internal/config"
)

// ScanDirectory walks the project directory and builds a Summary.
// It respects include/exclude lists from .hewd/config.yaml.
func ScanDirectory(root string) (*Summary, error) {
	cfg, cfgErr := config.Load(root)
	if cfgErr != nil {
		// Config errors are already structured by the config package.
		return nil, cfgErr
	}

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

	// Walk the repo
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			// We wrap this in a structured error so the user knows what failed.
			return cliutils.ErrHint(
				"failed to walk project directory",
				"check file permissions or exclude problematic paths using the config file",
			)
		}

		rel, relErr := filepath.Rel(root, path)
		if relErr != nil {
			return cliutils.Err("failed to compute relative path")
		}
		rel = filepath.ToSlash(rel)

		if rel == "." {
			return nil
		}

		// Always exclude .hewd/
		if strings.HasPrefix(rel, ".hewd/") {
			if entry.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Exclusion rule
		if shouldExclude(rel, excludes) {
			if entry.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Inclusion rule
		if len(includes) > 0 && !shouldInclude(rel, includes) {
			return nil
		}

		// Directory count
		if entry.IsDir() {
			s.Directories++
			return nil
		}

		// File count
		s.Files++

		name := entry.Name()

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
		// Errors returned from WalkDir callbacks are already structured,
		// but if WalkDir itself returns something unexpected, wrap it here.
		if he, ok := err.(cliutils.HewdError); ok {
			return nil, he
		}

		return nil, cliutils.ErrHint(
			"failed to scan project directory",
			"ensure all paths are readable or adjust include/exclude lists in .hewd/config.yaml",
		)
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
