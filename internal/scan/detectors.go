package scan

// This file declares detection maps for languages, documentation assets,
// and configuration assets. These maps are read by scanner.go,
// keeping all domain knowledge centralized and easy to maintain.

// -----------------------------------------------------------------------------
// Language Detection
// -----------------------------------------------------------------------------

// RealLanguages maps file extensions (without leading ".") to human-friendly
// language names. These values populate Summary.Languages.
var RealLanguages = map[string]string{
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

	// Scripting
	"sh":   "Shell",
	"bash": "Bash",
	"zsh":  "Zsh",
	"ps1":  "PowerShell",

	// Markup / data formats
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

// -----------------------------------------------------------------------------
// Documentation Detection
// -----------------------------------------------------------------------------

// DocumentationAssets maps filenames to human-friendly category labels.
// scanner.go will populate Summary.Documentation and Summary.DocsFound based
// on these entries.
var DocumentationAssets = map[string]string{
	"README.md":          "Project Overview",
	"CONTRIBUTING.md":    "Contribution Guide",
	"CHANGELOG.md":       "Changelog",
	"LICENSE":            "License File",
	"SECURITY.md":        "Security Policy",
	"CODE_OF_CONDUCT.md": "Code of Conduct",
}

// -----------------------------------------------------------------------------
// Configuration File Detection
// -----------------------------------------------------------------------------

// ConfigAssets maps filenames to classification strings. The scanner uses this
// to populate Summary.ConfigFiles.
var ConfigAssets = map[string]string{
	"go.mod":             "Go Module",
	"package.json":       "Node Package Manifest",
	"pyproject.toml":     "Python Project Config",
	"Dockerfile":         "Docker Build Config",
	"docker-compose.yml": "Docker Compose Config",
	"openapi.yaml":       "OpenAPI Specification",
	"openapi.yml":        "OpenAPI Specification",

	// Added after your bug report — ensures .gitignore is detected correctly:
	".gitignore": "Git Ignore File",
}
