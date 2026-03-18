package scan_test

import (
	"testing"

	"github.com/jbakchr/hewd/internal/scan"
	"github.com/jbakchr/hewd/internal/testutil"
)

// --------------------------------------------------
// Basic detection test
// --------------------------------------------------

func TestScan_DetectsLanguagesAndDocs(t *testing.T) {
	dir := testutil.CreateTempLayout(t, map[string]string{
		"README.md":  "hello",
		"main.go":    "package main",
		"pkg/app.js": "console.log('hi')",
	})

	s, err := scan.ScanDirectory(dir)
	if err != nil {
		t.Fatalf("scan error: %v", err)
	}

	if s.Languages["Go"] != 1 {
		t.Errorf("expected 1 Go file, got %d", s.Languages["Go"])
	}

	if s.Languages["JavaScript"] != 1 {
		t.Errorf("expected 1 JavaScript file, got %d", s.Languages["JavaScript"])
	}

	if !s.Documentation["README.md"] {
		t.Errorf("expected README.md to be detected")
	}
}

// --------------------------------------------------
// Exclude directories test
// --------------------------------------------------

func TestScan_ExcludeDirectory(t *testing.T) {
	dir := testutil.CreateTempLayout(t, map[string]string{
		".hewd/config.yaml": `
scan:
  exclude:
    - vendor
`,
		"main.go":          "package main",
		"vendor/ignore.go": "package ignored",
	})

	s, err := scan.ScanDirectory(dir)
	if err != nil {
		t.Fatalf("scan err: %v", err)
	}

	if s.Files != 1 { // vendor should be excluded
		t.Errorf("expected only 1 scanned file, got %d", s.Files)
	}
}
