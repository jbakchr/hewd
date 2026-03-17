package scan

// ───────────────────────────────────────────────────────────────
//
//	Language Detection
//
// ───────────────────────────────────────────────────────────────
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
var DocumentationAssets = map[string]string{
	"README.md":          "Project Overview",
	"CONTRIBUTING.md":    "Contribution Guide",
	"CHANGELOG.md":       "Changelog",
	"LICENSE":            "License File",
	"SECURITY.md":        "Security Policy",
	"CODE_OF_CONDUCT.md": "Code of Conduct",
}

// Map of config assets (file → label)
var ConfigAssets = map[string]string{
	"go.mod":             "Go Module",
	"package.json":       "Node Package Manifest",
	"pyproject.toml":     "Python Project Config",
	"Dockerfile":         "Docker Build Config",
	"docker-compose.yml": "Docker Compose Config",
	"openapi.yaml":       "OpenAPI Specification",
	"openapi.yml":        "OpenAPI Specification",
}
