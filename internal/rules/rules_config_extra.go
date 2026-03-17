package rules

import (
    "strings"
    "github.com/jbakchr/hewd/internal/scan"
)

func init() {
    RegisterRule(RuleHasDockerfileButNoCompose)
    RegisterRule(RuleMissingGitignore)
    RegisterRule(RuleCIFolderButNoWorkflows)
}

// 5. Dockerfile exists but no docker-compose.yml
func RuleHasDockerfileButNoCompose(s *scan.Summary) []Result {
    if len(s.ConfigFiles["Docker Build Config"]) == 0 {
        return nil
    }

    if len(s.ConfigFiles["Docker Compose Config"]) == 0 {
        return []Result{{
            ID:      "CFG_DOCKER_NO_COMPOSE",
            Level:   Info,
            Message: "Dockerfile found but no docker-compose.yml present.",
        }}
    }

    return nil
}

// 6. Missing .gitignore
func RuleMissingGitignore(s *scan.Summary) []Result {
    // Detect via ConfigFiles or direct scan
    found := false
    for _, files := range s.ConfigFiles {
        for _, f := range files {
            if strings.HasSuffix(f, ".gitignore") {
                found = true
            }
        }
    }

    if !found {
        return []Result{{
            ID:      "CFG_NO_GITIGNORE",
            Level:   Info,
            Message: ".gitignore file not found.",
        }}
    }

    return nil
}

// 7. .github/ folder exists but has no workflows
func RuleCIFolderButNoWorkflows(s *scan.Summary) []Result {
    hasGithubFolder := false
    hasWorkflows := false

    // Detect directories
    for _, files := range s.ConfigFiles {
        for _, f := range files {
            if strings.Contains(f, ".github") {
                hasGithubFolder = true
            }
            if strings.Contains(f, ".github/workflows") {
                hasWorkflows = true
            }
        }
    }

    if hasGithubFolder && !hasWorkflows {
        return []Result{{
            ID:      "CFG_GITHUB_NO_WORKFLOWS",
            Level:   Warn,
            Message: ".github/ folder found but no workflows detected.",
        }}
    }

    return nil
}