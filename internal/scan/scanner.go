package scan

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

//
// ───────────────────────────────────────────────────────────────
//   Scanner Implementation
// ───────────────────────────────────────────────────────────────
//
func ScanDirectory(root string) (*Summary, error) {
    s := &Summary{
        Languages:     map[string]int{},
        Documentation: map[string]bool{},
        DocsFound:     map[string][]string{},
        ConfigFiles:   map[string][]string{},
    }

    for name := range DocumentationAssets {
        s.Documentation[name] = false
    }

    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil { return err }
        if path == root { return nil }

        if info.IsDir() {
            s.Directories++
            return nil
        }

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
        ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(name)), ".")
        if lang, ok := RealLanguages[ext]; ok {
            s.Languages[lang]++
        }

        return nil
    })

    if err != nil {
        return nil, fmt.Errorf("error scanning: %w", err)
    }

    return s, nil
}