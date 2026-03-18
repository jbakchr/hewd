package rules

import (
	"strings"

	"github.com/jbakchr/hewd/internal/config"
	"github.com/jbakchr/hewd/internal/scan"
)

func RunAll(summary *scan.Summary, cfg *config.Config, onlyCats, exceptCats []string) []Result {

	results := []Result{}

	// Rule disabling (from config.rules: map[string]bool)
	disabled := map[string]bool{}
	if cfg != nil && cfg.Rules != nil {
		for ruleID, enabled := range cfg.Rules {
			if !enabled {
				disabled[ruleID] = true
			}
		}
	}

	// Category filtering maps
	only := make(map[string]bool)
	for _, c := range onlyCats {
		only[strings.ToLower(c)] = true
	}

	except := make(map[string]bool)
	for _, c := range exceptCats {
		except[strings.ToLower(c)] = true
	}

	for _, reg := range allRules {

		cat := strings.ToLower(reg.Category)

		// category filtering: only vs except
		if len(only) > 0 && !only[cat] {
			continue
		}
		if except[cat] {
			continue
		}

		// rule disabling
		if disabled[reg.ID] {
			continue
		}

		out := reg.Func(summary)

		// severity override from config.weights
		if cfg != nil && cfg.Weights != nil {
			for i := range out {
				if override, ok := cfg.Weights[out[i].ID]; ok {
					out[i].Level = LevelFromInt(override)
				}
			}
		}

		results = append(results, out...)
	}

	return results
}
