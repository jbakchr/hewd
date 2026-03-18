package testutil

import (
	"os"
	"path/filepath"
	"testing"
)

// CreateTempLayout creates a temporary directory with the given file structure.
// `files` maps paths -> content. Directories are created automatically.
func CreateTempLayout(t *testing.T, files map[string]string) string {
	t.Helper()

	dir := t.TempDir()

	for path, content := range files {
		full := filepath.Join(dir, path)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatalf("failed to mkdir: %v", err)
		}
		if err := os.WriteFile(full, []byte(content), 0644); err != nil {
			t.Fatalf("failed to write file: %v", err)
		}
	}

	return dir
}
