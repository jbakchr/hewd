package helptext

const ScanUse = "scan"

const ScanShort = "Scan the repository and detect documentation, config, and structure indicators."

const ScanLong = `hewd scan performs a fast analysis of the current repository. It detects
documentation files, configuration files, programming languages, metadata, and
structural indicators such as docs/ directories or CI workflows.

Features:

  • Fast repository scanning
  • Language, file, and metadata detection
  • Pretty, JSON, YAML, or Markdown output
  • CI-friendly machine-readable modes

Use hewd scan to get a quick overview of project structure before running
'hewd doctor' or generating reports for automation.`

const ScanExample = `
  # Pretty scan output
  hewd scan --pretty

  # Output JSON
  hewd scan --json

  # Output YAML
  hewd scan --yaml

  # Save JSON output
  hewd scan --json > scan.json

  # Combine scan and diagnostics
  hewd scan --pretty && hewd doctor
`
