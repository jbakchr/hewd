package helptext

const ExportUse = "export"

const ExportShort = "Export a complete machine-readable project health report."

const ExportLong = `hewd export generates a machine-readable health report for the current
repository. The exported data uses hewd's stable MachineOutput schema and
includes scores, rule results, fixable items, metadata, and timestamps.

Features:

  • Full project health report
  • JSON or YAML export
  • CI-friendly output structure
  • Compatible with the diff engine

Use hewd export to produce reports for dashboards, CI pipelines, automation, or
as input for hewd diff when comparing changes over time.`

const ExportExample = `
  # Export project health to JSON
  hewd export --output hewd.json

  # Export YAML
  hewd export --yaml --output hewd.yaml

  # Pretty JSON to stdout
  hewd export --json --pretty

  # Generate old/new reports for diff
  hewd export --output old.json
  # make changes...
  hewd export --output new.json
  hewd diff old.json new.json

  # Pipe JSON into another tool
  hewd export --json | jq '.score'
`
