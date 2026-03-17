package config

import (
    "os"
    "path/filepath"

    "gopkg.in/yaml.v3"
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

    _, err := os.Stat(cfgPath)
    if os.IsNotExist(err) {
        // No config; return defaults
        return &Config{
            Rules:   map[string]bool{},
            Weights: map[string]int{},
        }, nil
    }

    data, err := os.ReadFile(cfgPath)
    if err != nil {
        return nil, err
    }

    var cfg Config
    if err := yaml.Unmarshal(data, &cfg); err != nil {
        return nil, err
    }

    return &cfg, nil
}