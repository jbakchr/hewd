package helptext

import "github.com/jbakchr/hewd/internal/cliutils"

const ScanUse = "scan"

const ScanShort = "Scan the repository and detect documentation, config, and structure indicators."

// Tagline shown at the very top of `hewd scan` output.
const ScanLongTagline = cliutils.CyanBold +
	"hewd " + cliutils.WhiteBoldItalic + "– scan" +
	cliutils.Reset + "\n\n"

const ScanLong = ScanLongTagline + cliutils.Cyan + "hewd " + cliutils.Reset +
	`scan performs a fast analysis of the current repository. It detects
documentation files, configuration files, programming languages, metadata, and
structural indicators such as docs/ directories or CI workflows.

` + cliutils.WhiteBold + `Features:` + cliutils.Reset + `

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
