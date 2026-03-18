package scan

// Summary holds the aggregated results of scanning a project directory.
type Summary struct {
	// Basic stats
	Files       int `json:"files" yaml:"files"`
	Directories int `json:"directories" yaml:"directories"`

	// Language → file count
	Languages map[string]int `json:"languages" yaml:"languages"`

	// Documentation presence map: "README.md": true/false
	Documentation map[string]bool `json:"documentation" yaml:"documentation"`

	// DocsFound maps a doc type → list of file paths
	DocsFound map[string][]string `json:"documentation_files" yaml:"documentation_files"`

	// ConfigFiles maps config type → paths
	ConfigFiles map[string][]string `json:"config_files" yaml:"config_files"`
}
