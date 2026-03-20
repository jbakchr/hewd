package helptext

const ScanUse = "scan"

const ScanShort = "Scan the repository and detect documentation, config, languages, and structure indicators."

const ScanLong = `hewd scan performs a fast, lightweight analysis of the current repository.
It detects documentation files, configuration files, programming languages,
project metadata, and structural indicators such as the presence of a docs/
directory or CI workflows.

This command provides a high-level overview of your project's documentation and
structure. It is often the first step before running 'hewd doctor' or exporting
a machine-readable health report.

Scan output supports multiple formats:

  • Pretty (human-readable)
  • JSON
  • YAML

This command is safe to run in both local development and CI environments.`

const ScanExample = `
  # Scan the current repository (pretty output)
  hewd scan --pretty

  # Scan and output JSON
  hewd scan --json

  # Scan and output YAML
  hewd scan --yaml

  # Save JSON output
  hewd scan --json > scan.json

  # Combine scan + doctor
  hewd scan --pretty && hewd doctor
`
