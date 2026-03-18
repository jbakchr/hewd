package rules

import (
    "github.com/jbakchr/hewd/internal/config"
    "github.com/jbakchr/hewd/internal/scan"
)

// RunAll executes all registered rules against the scan.Summary,
// applying config-driven rule disabling and severity overrides.
func RunAll(summary *scan.Summary, cfg *config.Config) []Result {
    results := []Result{}

    // Build disabled rule map
    disabled := map[string]bool{}
    if cfg != nil && cfg.Rules != nil {
        for ruleID, enabled := range cfg.Rules {
            // rules: { RULE_ID: false } means disabled
            if !enabled {
                disabled[ruleID] = true
            }
        }
    }

    // Run each registered rule
    for _, reg := range allRules {
        // Skip disabled rules
        if disabled[reg.ID] {
            continue
        }

        // Evaluate rule
        out := reg.Func(summary)

        // Apply config severity overrides
        if cfg != nil && cfg.Weights != nil {
            for i := range out {
                if newWeight, ok := cfg.Weights[out[i].ID]; ok {
                    out[i].Level = LevelFromInt(newWeight)
                }
            }
        }

        results = append(results, out...)
    }

    return results
}