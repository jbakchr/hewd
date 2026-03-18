package scan

// Summary contains all information collected by the scan engine.
// It is the core data structure passed into the rules engine.
type Summary struct {
	// Basic statistics
	Files       int `json:"files" yaml:"files"`
	Directories int `json:"directories" yaml:"directories"`

	// Language -> file count
	Languages map[string]int `json:"languages" yaml:"languages"`

	// Documentation presence map: "README.md": true/false
	Documentation map[string]bool `json:"documentation" yaml:"documentation"`

	// DocsFound maps "Project Overview" -> []paths
	DocsFound map[string][]string `json:"documentation_files" yaml:"documentation_files"`

	// ConfigFiles maps "Go Module" -> []paths
	ConfigFiles map[string][]string `json:"config_files" yaml:"config_files"`
}
