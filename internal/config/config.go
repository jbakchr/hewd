package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/jbakchr/hewd/internal/cliutils"
)

// Config holds all configurable options for hewd.
type Config struct {
	Rules   map[string]bool `yaml:"rules"`
	Weights map[string]int  `yaml:"weights"`

	Scan struct {
		Include []string `yaml:"include"`
		Exclude []string `yaml:"exclude"`
	} `yaml:"scan"`
}

// Load loads `.hewd/config.yaml` if present.
// If the file does not exist, it returns an empty config and no error.
func Load(root string) (*Config, error) {
	cfgPath := filepath.Join(root, ".hewd", "config.yaml")

	_, statErr := os.Stat(cfgPath)
	if os.IsNotExist(statErr) {
		// No config; return defaults
		return &Config{
			Rules:   map[string]bool{},
			Weights: map[string]int{},
		}, nil
	}

	if statErr != nil {
		return nil, cliutils.ErrHint(
			"failed to stat config file",
			"check file permissions and ensure the path is readable",
		)
	}

	data, readErr := os.ReadFile(cfgPath)
	if readErr != nil {
		return nil, cliutils.ErrHint(
			"failed to read config file",
			"ensure the file exists and has correct permissions",
		)
	}

	var cfg Config
	if unmarshalErr := yaml.Unmarshal(data, &cfg); unmarshalErr != nil {
		return nil, cliutils.ErrHint(
			"failed to parse config",
			"ensure .hewd/config.yaml contains valid yaml",
		)
	}

	return &cfg, nil
}
