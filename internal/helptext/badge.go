package helptext

const BadgeUse = "badge"

const BadgeShort = "Generate an SVG badge displaying the project's overall health score."

const BadgeLong = `hewd badge generates a standalone SVG badge representing the project's overall
health score. The badge is similar to common README badges and uses color
indicators to show project health at a glance.

Badge creation is fully local—ideal for CI pipelines and documentation.`

const BadgeExample = `
  # Create a badge
  hewd badge --output badge.svg

  # Save badge in docs/
  hewd badge --output docs/health-badge.svg

  # Regenerate after diagnostics
  hewd doctor --json > report.json
  hewd badge --output badge.svg
`
