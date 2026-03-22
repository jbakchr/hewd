package helptext

import "github.com/jbakchr/hewd/internal/cliutils"

const RootUse = "hewd"

const RootShort = "Analyze, score, and improve the health of software repositories."

// Tagline shown at the very top of `hewd` output.
const RootLongTagline = cliutils.CyanBold +
	"hewd " + cliutils.WhiteBoldItalic + "– repository health diagnostics, scoring, and automated fixes" +
	cliutils.Reset + "\n\n"

// Full description.
const RootLong = RootLongTagline +
	cliutils.Cyan + "hewd" + cliutils.Reset +
	` analyzes repository health by evaluating documentation, configuration,
and structural conventions. It provides fast scanning, actionable feedback,
health scores, diff reports, and automated fixes.

` + cliutils.WhiteBold + `Features:` + cliutils.Reset + `

  • Fast, dependency-free repository scanner
  • Curated rules for documentation, structure, and configuration
  • Automated fixes for common issues
  • JSON, YAML, Markdown, and pretty outputs
  • Regression gating for CI pipelines
  • GitHub Action for PR comments
  • SVG badge generation

Use ` + cliutils.Cyan + `hewd` + cliutils.Reset + ` to maintain consistent documentation, detect regressions, enforce
standards, and track repository maturity over time.`

// Examples shown in `hewd --help`.
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
  hewd init`
