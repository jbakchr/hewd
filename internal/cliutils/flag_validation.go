package cliutils

import "fmt"

// ValidateOutputFormatFlags enforces universal output flag rules:
//
// - Only one of: --json, --yaml, --md
// - --pretty applies only to JSON
// - Returns an error if flags are in conflict
//
// Pass boolean values for each flag and the name of the command
// (used for clearer error messages).
func ValidateOutputFormatFlags(jsonOut, yamlOut, mdOut, pretty bool, commandName string) error {

	// Only one machine format at a time
	if jsonOut && yamlOut {
		return fmt.Errorf("%s: cannot combine --json and --yaml", commandName)
	}
	if jsonOut && mdOut {
		return fmt.Errorf("%s: cannot combine --json and --md", commandName)
	}
	if yamlOut && mdOut {
		return fmt.Errorf("%s: cannot combine --yaml and --md", commandName)
	}

	// Pretty only applies to JSON
	if pretty && yamlOut {
		return fmt.Errorf("%s: cannot combine --yaml and --pretty (pretty applies only to JSON)", commandName)
	}
	if pretty && mdOut {
		return fmt.Errorf("%s: cannot combine --md and --pretty (pretty applies only to JSON)", commandName)
	}

	return nil
}
