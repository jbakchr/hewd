package rules

import (
	"strings"

	"github.com/jbakchr/hewd/internal/scan"
)

// This file contains additional configuration and CI-related rules.
// These rules cover areas such as git hygiene, CI workflow existence,
// and containerization best practices.

func init() {
	RegisterRule("CFG_DOCKER_NO_COMPOSE", RuleHasDockerfileButNoCompose)
	RegisterRule("CFG_NO_GITIGNORE", RuleMissingGitignore)
	RegisterRule("CFG_GITHUB_NO_WORKFLOWS", RuleCIFolderButNoWorkflows)
}

// -----------------------------------------------------------------------------
// 1. Dockerfile present but no docker-compose.yml
// -----------------------------------------------------------------------------

func RuleHasDockerfileButNoCompose(s interface{}) []Result {
	summary := s.(*scan.Summary)

	hasDockerfile := len(summary.ConfigFiles["Docker Build Config"]) > 0
	hasCompose := len(summary.ConfigFiles["Docker Compose Config"]) > 0

	if !hasDockerfile {
		return nil
	}

	if !hasCompose {
		return []Result{{
			ID:      "CFG_DOCKER_NO_COMPOSE",
			Level:   Info,
			Message: "Dockerfile found but no docker-compose.yml present.",
			File:    first(summary.ConfigFiles["Docker Build Config"]),
		}}
	}

	return nil
}

// -----------------------------------------------------------------------------
// 2. .gitignore should exist in most projects
// -----------------------------------------------------------------------------

func RuleMissingGitignore(s interface{}) []Result {
	summary := s.(*scan.Summary)

	// Detect via ConfigFiles
	if len(summary.ConfigFiles["Git Ignore File"]) > 0 {
		return nil
	}

	return []Result{{
		ID:      "CFG_NO_GITIGNORE",
		Level:   Info,
		Message: ".gitignore file not found.",
	}}
}

// -----------------------------------------------------------------------------
// 3. .github folder present but no workflow files
// -----------------------------------------------------------------------------

func RuleCIFolderButNoWorkflows(s interface{}) []Result {
	summary := s.(*scan.Summary)

	hasGithubFolder := false
	hasWorkflows := false

	for cfgType, files := range summary.ConfigFiles {
		for _, f := range files {
			if strings.Contains(f, ".github/") {
				hasGithubFolder = true
			}
			if cfgType == ".github/workflows" ||
				strings.Contains(f, ".github/workflows/") {
				hasWorkflows = true
			}
		}
	}

	if hasGithubFolder && !hasWorkflows {
		return []Result{{
			ID:      "CFG_GITHUB_NO_WORKFLOWS",
			Level:   Warn,
			Message: ".github/ folder found but no workflow files detected.",
		}}
	}

	return nil
}
