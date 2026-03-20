package helptext

const FixUse = "fix"

const FixShort = "Automatically generate missing documentation, structure, and CI files."

const FixLong = `hewd fix analyzes the repository for missing documentation, structure, and
configuration files, and automatically generates recommended files.

By default, fixes are shown as a dry-run. Use --apply to write changes to disk.
Existing files are never overwritten.`

const FixExample = `
  # Dry-run (show fixes)
  hewd fix

  # Apply fixes
  hewd fix --apply

  # Combine with diagnostics
  hewd fix --apply && hewd doctor

  # Save preview to a file
  hewd fix > fix-preview.txt
`
