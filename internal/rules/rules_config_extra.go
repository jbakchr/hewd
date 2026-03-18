package rules

import (
	"strings"

	"github.com/jbakchr/hewd/internal/scan"
)

func init() {
	RegisterRule("CFG_DOCKER_NO_COMPOSE", "config", RuleHasDockerfileButNoCompose)
	RegisterRule("CFG_NO_GITIGNORE", "config", RuleMissingGitignore)
	RegisterRule("CFG_GITHUB_NO_WORKFLOWS", "config", RuleCIFolderButNoWorkflows)
}

// 1. Dockerfile present but no docker-compose.yml
func RuleHasDockerfileButNoCompose(s interface{}) []Result {
	sum := s.(*scan.Summary)

	hasDocker := len(sum.ConfigFiles["Docker Build Config"]) > 0
	hasCompose := len(sum.ConfigFiles["Docker Compose Config"]) > 0

	if !hasDocker {
		return nil
	}

	if !hasCompose {
		return []Result{{
			ID:      "CFG_DOCKER_NO_COMPOSE",
			Level:   Info,
			Message: "Dockerfile found but no docker-compose.yml present.",
			File:    first(sum.ConfigFiles["Docker Build Config"]),
		}}
	}

	return nil
}

// 2. Missing .gitignore
func RuleMissingGitignore(s interface{}) []Result {
	sum := s.(*scan.Summary)

	if len(sum.ConfigFiles["Git Ignore File"]) > 0 {
		return nil
	}

	return []Result{{
		ID:      "CFG_NO_GITIGNORE",
		Level:   Info,
		Message: ".gitignore file not found.",
	}}
}

// 3. .github folder exists but no workflow files
func RuleCIFolderButNoWorkflows(s interface{}) []Result {
	sum := s.(*scan.Summary)

	hasGithub := false
	hasWorkflows := false

	for _, files := range sum.ConfigFiles {
		for _, f := range files {
			if strings.Contains(f, ".github/") {
				hasGithub = true
			}
			if strings.Contains(f, ".github/workflows/") {
				hasWorkflows = true
			}
		}
	}

	if hasGithub && !hasWorkflows {
		return []Result{{
			ID:      "CFG_GITHUB_NO_WORKFLOWS",
			Level:   Warn,
			Message: ".github/ folder found but no workflow files detected.",
		}}
	}

	return nil
}
