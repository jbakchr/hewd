package helptext

const DoctorUse = "doctor"

const DoctorShort = "Run full diagnostics and compute documentation, config, and structure scores."

const DoctorLong = `hewd doctor runs the full diagnostic engine on the current repository. It
evaluates documentation, configuration, and structure using a curated set of
rules, each with a severity level (info, warn, error).

The doctor command produces:

  • Category scores (documentation, config, structure)
  • An overall project health score
  • Detailed issue reports with severity metadata
  • A list of fixable problems
  • Optional Markdown, JSON, or YAML output

Markdown output is ideal for pull request comments, while JSON and YAML are
well‑suited for CI pipelines and automated quality gates.

Use 'hewd doctor' regularly to enforce documentation standards, detect
regressions, and maintain consistent quality across repositories.`

const DoctorExample = `
  # Full diagnostics (pretty output)
  hewd doctor

  # Markdown output (ideal for PR comments)
  hewd doctor --md > health.md

  # JSON for CI
  hewd doctor --json > doctor.json

  # YAML output
  hewd doctor --yaml

  # Only documentation rules
  hewd doctor --only documentation

  # Exclude config rules
  hewd doctor --except config

  # Fail CI on warning-level issues
  hewd doctor --fail-on=warn
`
