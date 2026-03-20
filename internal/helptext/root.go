package helptext

const RootUse = "hewd"

const RootShort = "Analyze, score, and improve the health of software repositories."

const RootLong = `hewd analyzes the health of software repositories by evaluating documentation,
configuration, and structural conventions. It provides actionable feedback,
health scores, diff reports, and automated fixes.

Features:

  • Fast, dependency‑free repository scanner
  • Curated rule engine for documentation, structure, and config
  • Automated fix mode for common issues
  • JSON, YAML, Markdown, and pretty outputs
  • Regression gating for CI pipelines
  • GitHub Action with PR comment updates
  • SVG badge generation

Use hewd to maintain consistent documentation, detect regressions, enforce
standards, and track repository maturity over time.`

const RootExample = `
  # Scan the repository and show a high-level summary
  hewd scan --pretty

  # Run full diagnostics and generate a Markdown report
  hewd doctor --md > health.md

  # Export machine-readable project health
  hewd export --output hewd.json

  # Compare two reports (detect new/resolved issues)
  hewd diff old.json new.json --md > diff.md

  # Automatically fix missing documentation or CI files
  hewd fix --apply

  # Generate an SVG badge showing project health score
  hewd badge --output badge.svg

  # Initialize a .hewd/config.yaml configuration file
  hewd init
`
