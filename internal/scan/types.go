package scan

type Summary struct {
	Files         int                 `json:"files" yaml:"files"`
	Directories   int                 `json:"directories" yaml:"directories"`
	Languages     map[string]int      `json:"languages" yaml:"languages"`
	Documentation map[string]bool     `json:"documentation" yaml:"documentation"`
	DocsFound     map[string][]string `json:"documentation_files" yaml:"documentation_files"`
	ConfigFiles   map[string][]string `json:"config_files" yaml:"config_files"`
}
