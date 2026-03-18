package rules

import "github.com/jbakchr/hewd/internal/scan"

func init() {
	RegisterRule("CFG_GOMOD_NO_LICENSE", "config", RuleGoModWithoutLicense)
	RegisterRule("CFG_NODE_NO_LOCKFILE", "config", RuleNodePackageWithoutLockfile)
}

// 1. go.mod exists but LICENSE missing
func RuleGoModWithoutLicense(s interface{}) []Result {
	sum := s.(*scan.Summary)

	if len(sum.ConfigFiles["Go Module"]) == 0 {
		return nil
	}

	if sum.Documentation["LICENSE"] {
		return nil
	}

	return []Result{{
		ID:      "CFG_GOMOD_NO_LICENSE",
		Level:   Warn,
		Message: "go.mod is present but the project does not include a LICENSE file.",
		File:    first(sum.ConfigFiles["Go Module"]),
	}}
}

// 2. package.json exists but no lockfile
func RuleNodePackageWithoutLockfile(s interface{}) []Result {
	sum := s.(*scan.Summary)

	if len(sum.ConfigFiles["Node Package Manifest"]) == 0 {
		return nil
	}

	hasNpmLock := len(sum.ConfigFiles["package-lock.json"]) > 0
	hasYarnLock := len(sum.ConfigFiles["yarn.lock"]) > 0

	if hasNpmLock || hasYarnLock {
		return nil
	}

	return []Result{{
		ID:      "CFG_NODE_NO_LOCKFILE",
		Level:   Warn,
		Message: "package.json is present but no lockfile detected (package-lock.json or yarn.lock).",
		File:    first(sum.ConfigFiles["Node Package Manifest"]),
	}}
}
