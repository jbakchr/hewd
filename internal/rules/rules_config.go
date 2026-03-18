package rules

import (
	"github.com/jbakchr/hewd/internal/scan"
)

// This file contains the *core* configuration-related rules.
// Additional, more advanced config rules are implemented in
// rules_config_extra.go.

func init() {
	RegisterRule("CFG_GOMOD_NO_LICENSE", RuleGoModWithoutLicense)
	RegisterRule("CFG_NODE_NO_LOCKFILE", RuleNodePackageWithoutLockfile)
}

// -----------------------------------------------------------------------------
// 1. Go project should have LICENSE when go.mod exists
// -----------------------------------------------------------------------------

func RuleGoModWithoutLicense(s interface{}) []Result {
	summary := s.(*scan.Summary)

	// If there's no go.mod, the rule does not apply
	if len(summary.ConfigFiles["Go Module"]) == 0 {
		return nil
	}

	// If LICENSE exists, pass
	if summary.Documentation["LICENSE"] {
		return nil
	}

	return []Result{{
		ID:      "CFG_GOMOD_NO_LICENSE",
		Level:   Warn,
		Message: "go.mod is present but the project does not include a LICENSE file.",
		File:    first(summary.ConfigFiles["Go Module"]),
	}}
}

// -----------------------------------------------------------------------------
// 2. Node project should have a lockfile when package.json exists
// -----------------------------------------------------------------------------

func RuleNodePackageWithoutLockfile(s interface{}) []Result {
	summary := s.(*scan.Summary)

	// Not a Node project
	if len(summary.ConfigFiles["Node Package Manifest"]) == 0 {
		return nil
	}

	// Node lockfiles to check
	hasNpmLock := len(summary.ConfigFiles["package-lock.json"]) > 0
	hasYarnLock := len(summary.ConfigFiles["yarn.lock"]) > 0

	if hasNpmLock || hasYarnLock {
		return nil
	}

	return []Result{{
		ID:      "CFG_NODE_NO_LOCKFILE",
		Level:   Warn,
		Message: "package.json is present but no lockfile detected (package-lock.json or yarn.lock).",
		File:    first(summary.ConfigFiles["Node Package Manifest"]),
	}}
}
