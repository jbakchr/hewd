package helptext

const ExportUse = "export"

const ExportShort = "Export a complete machine-readable project health report."

const ExportLong = `hewd export generates a complete, machine-readable report describing the
documentation, configuration, and structural health of the current repository.

The exported report follows hewd's stable MachineOutput schema, which includes:

  • Overall health score
  • Category scores
  • Rule results with severity levels
  • Fixable issues
  • Version metadata
  • Timestamps

Exports are ideal for CI pipelines, dashboards, trend tracking, and diff
comparisons.`

const ExportExample = `
  # Export project health to JSON
  hewd export --output hewd.json

  # Export to YAML
  hewd export --yaml --output hewd.yaml

  # Pretty JSON to stdout
  hewd export --json --pretty

  # Generate before/after reports for diff
  hewd export --output old.json
  # make changes...
  hewd export --output new.json
  hewd diff old.json new.json

  # Pipe JSON output
  hewd export --json | jq '.score'
`
