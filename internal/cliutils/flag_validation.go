package cliutils

// ValidateOutputFormatFlags enforces universal output-format flag rules.
//
// Rules:
//   - only one of --json, --yaml, or --md may be active
//   - --pretty applies only to --json
//   - every violation produces a HewdError with an actionable hint
//
// Parameters:
//
//	jsonOut  = flag status
//	yamlOut  = flag status
//	mdOut    = flag status
//	pretty   = flag status
//	cmdName  = the subcommand invoking validation (e.g., "scan", "doctor")
//
// Returns HewdError if the flags are invalid.
func ValidateOutputFormatFlags(jsonOut, yamlOut, mdOut, pretty bool, cmdName string) error {

	// -------------------------------------------------------------------------
	// Mutually exclusive: --json AND --yaml
	// -------------------------------------------------------------------------
	if jsonOut && yamlOut {
		return ErrHint(
			"cannot combine --json and --yaml",
			"use only one machine-readable format at a time",
		)
	}

	// -------------------------------------------------------------------------
	// Mutually exclusive: --json AND --md
	// -------------------------------------------------------------------------
	if jsonOut && mdOut {
		return ErrHint(
			"cannot combine --json and --md",
			"use only one machine-readable format at a time",
		)
	}

	// -------------------------------------------------------------------------
	// Mutually exclusive: --yaml AND --md
	// -------------------------------------------------------------------------
	if yamlOut && mdOut {
		return ErrHint(
			"cannot combine --yaml and --md",
			"use only one machine-readable format at a time",
		)
	}

	// -------------------------------------------------------------------------
	// pretty + yaml is invalid
	// -------------------------------------------------------------------------
	if pretty && yamlOut {
		return ErrHint(
			"cannot combine --yaml and --pretty",
			"the --pretty flag only applies to JSON output",
		)
	}

	// -------------------------------------------------------------------------
	// pretty + md is invalid
	// -------------------------------------------------------------------------
	if pretty && mdOut {
		return ErrHint(
			"cannot combine --md and --pretty",
			"the --pretty flag only applies to JSON output",
		)
	}

	// If no conflicts were found:
	return nil
}
