package helptext

const RootUse = "hewd"

const RootShort = "Analyze, score, and improve the health of software repositories."

const RootLong = `hewd analyzes the health of software repositories by evaluating documentation,
configuration, and structural conventions. It provides actionable feedback,
health scores, diff reports, and automated fixes.

Features:

  • Fast, dependency-free repository scanner
  • Curated rules for documentation, structure, and configuration
  • Automated fixes for common issues
  • JSON, YAML, Markdown, and pretty outputs
  • Regression gating for CI pipelines
  • GitHub Action for PR comments
  • SVG badge generation

Use hewd to maintain consistent documentation, detect regressions, enforce
standards, and track repository maturity over time.`

const RootExample = `
  # Scan a repository
  hewd scan --pretty

  # Run full diagnostics and generate Markdown
  hewd doctor --md > health.md

  # Export machine-readable project health
  hewd export --output hewd.json

  # Compare reports using the diff engine
  hewd diff old.json new.json --md > diff.md

  # Apply automated fixes
  hewd fix --apply

  # Generate an SVG badge
  hewd badge --output badge.svg

  # Initialize a hewd configuration file
  hewd init
`
