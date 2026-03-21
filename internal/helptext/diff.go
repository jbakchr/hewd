package helptext

const DiffUse = "diff <old.json> <new.json>"

const DiffShort = "Compare two hewd reports and show score, category, and issue differences."

const DiffLong = `hewd diff compares two machine-readable hewd reports and highlights how a
project's health has changed. It computes score deltas, category changes, new
and resolved issues, and grouped summaries.

Features:

  • Score and category deltas
  • New and resolved issue detection
  • Pretty, JSON, YAML, and Markdown output
  • Regression gating for CI pipelines
  • GitHub-friendly Markdown formatting

Use hewd diff to track repository trends, enforce quality standards, and gate
pull requests through automated checks.`

const DiffExample = `
  # Pretty diff
  hewd diff old.json new.json

  # Markdown diff for PRs
  hewd diff old.json new.json --md > diff.md

  # JSON diff for CI
  hewd diff old.json new.json --json > diff.json

  # YAML diff
  hewd diff old.json new.json --yaml

  # Fail if score drops by 5 points or more
  hewd diff old.json new.json --fail-on-score-drop=5

  # Fail if new error-level issues appear
  hewd diff old.json new.json --fail-on-new-errors

  # Fail on any regression
  hewd diff old.json new.json --fail-on-any-regression
`
