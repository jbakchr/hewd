package helptext

const BadgeUse = "badge"

const BadgeShort = "Generate an SVG badge displaying the project's health score."

const BadgeLong = `hewd badge generates a standalone SVG badge representing the project's overall
health score. The badge uses color indicators to communicate repository health
at a glance and can be embedded in documentation or dashboards.

Features:

  • Local SVG generation (no external services)
  • Consistent, readable badge styling
  • Ideal for README badges and CI artifacts

Use hewd badge to visualize your project's health score or publish status
badges for documentation and automation workflows.`

const BadgeExample = `
  # Generate an SVG badge
  hewd badge --output badge.svg

  # Save badge into a documentation folder
  hewd badge --output docs/health-badge.svg

  # Regenerate badge after diagnostics
  hewd doctor --json > report.json
  hewd badge --output badge.svg
`
