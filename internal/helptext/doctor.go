package helptext

const DoctorUse = "doctor"

const DoctorShort = "Run full diagnostics and compute documentation, config, and structure scores."

const DoctorLong = `hewd doctor evaluates repository documentation, configuration, and structure
using a curated rule engine. Each rule has a severity level (info, warn, error)
and contributes to category and overall health scores.

Features:

  • Detailed rule evaluation with severity levels
  • Category scores (documentation, config, structure)
  • Overall project health score
  • Fixable item detection
  • Pretty, JSON, YAML, or Markdown output
  • CI integration via fail-on flags

Use hewd doctor to verify project health, enforce standards, and detect
regressions across repositories.`

const DoctorExample = `
  # Run full diagnostics
  hewd doctor

  # Output Markdown report
  hewd doctor --md > health.md

  # Export JSON for CI
  hewd doctor --json > doctor.json

  # Output YAML
  hewd doctor --yaml

  # Run only documentation rules
  hewd doctor --only documentation

  # Exclude config rules
  hewd doctor --except config

  # Fail CI on warnings or worse
  hewd doctor --fail-on=warn
`
