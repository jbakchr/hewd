package rules

import "github.com/jbakchr/hewd/internal/scan"

func init() {
    RegisterRule(RuleGoModWithoutLicense)
    RegisterRule(RuleNodePackageWithoutLockfile)
}

func RuleGoModWithoutLicense(s *scan.Summary) []Result {
    if len(s.ConfigFiles["Go Module"]) == 0 {
        return nil
    }
    if s.Documentation["LICENSE"] {
        return nil
    }
    return []Result{{
        ID:      "CFG_GOMOD_NO_LICENSE",
        Level:   Warn,
        Message: "go.mod present but no LICENSE file found.",
        File:    s.ConfigFiles["Go Module"][0],
    }}
}

func RuleNodePackageWithoutLockfile(s *scan.Summary) []Result {
    if len(s.ConfigFiles["Node Package Manifest"]) == 0 {
        return nil
    }
    // detect missing lockfile
    if len(s.ConfigFiles["package-lock.json"]) == 0 && len(s.ConfigFiles["yarn.lock"]) == 0 {
        return []Result{{
            ID:      "CFG_NODE_NO_LOCKFILE",
            Level:   Warn,
            Message: "package.json present but no lockfile detected.",
        }}
    }
    return nil
}