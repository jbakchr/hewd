package helptext

const FixUse = "fix"

const FixShort = "Automatically generate missing documentation, structure, and CI files."

const FixLong = `hewd fix analyzes the repository for missing documentation, structure, and
configuration files, then generates appropriate recommended files.

Features:

  • Generates missing documentation files
  • Creates CI workflow templates
  • Adds missing project structure directories
  • Safe repeated runs (never overwrites files)

Use hewd fix to bootstrap documentation, enforce standards, and prepare a
repository for diagnostics or publication.`

const FixExample = `
  # Show fixable issues (dry-run)
  hewd fix

  # Apply fixes and write files
  hewd fix --apply

  # Apply fixes before running diagnostics
  hewd fix --apply && hewd doctor

  # Save preview to a file
  hewd fix > fix-preview.txt
`
