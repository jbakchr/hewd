package doctor

import (
    cfgpkg "github.com/jbakchr/hewd/internal/config"
    "github.com/jbakchr/hewd/internal/scan"
)

func RunAll(s *scan.Summary, cfg *cfgpkg.Config) Result {
    result := Result{}

    total := 0
    max := 0

    for _, rule := range Rules {

        // 1. Rule disable logic
        if enabled, ok := cfg.Rules[rule.ID]; ok && !enabled {
            continue
        }

        // 2. Determine weight (config overrides default)
        w := rule.Weight
        if override, ok := cfg.Weights[rule.ID]; ok {
            w = override
        }
        if w <= 0 {
            w = 1
        }

        passed, message := rule.Check(s)

        result.Findings = append(result.Findings, Finding{
            RuleID:  rule.ID,
            Passed:  passed,
            Message: message,
        })

        max += w
        if passed {
            total += w
        }
    }

    result.Score = total
    result.MaxScore = max

    return result
}