package helptext

const DiffUse = "diff <old.json> <new.json>"

const DiffShort = "Compare two hewd JSON reports and show score, category, and issue differences."

const DiffLong = `hewd diff compares two machine-readable hewd reports (JSON) and highlights how
a project's health has changed. The diff engine computes:

  • Score deltas
  • Category score changes
  • New issues
  • Resolved issues
  • Grouped and sorted issue summaries

Supports pretty, JSON, YAML, and Markdown output. Regression gating options allow
CI pipelines to fail automatically when regressions occur.`

const DiffExample = `
  # Pretty diff
  hewd diff old.json new.json

  # Markdown diff (ideal for PRs)
  hewd diff old.json new.json --md > diff.md

  # JSON diff for CI
  hewd diff old.json new.json --json > diff.json

  # YAML diff
  hewd diff old.json new.json --yaml

  # Fail if score drops by 5+
  hewd diff old.json new.json --fail-on-score-drop=5

  # Fail on new error-level issues
  hewd diff old.json new.json --fail-on-new-errors

  # Fail on any regression
  hewd diff old.json new.json --fail-on-any-regression
`
