package helptext

import (
	"strings"

	"github.com/jbakchr/hewd/internal/formatter"
)

var featuresList = []string{
	"• Fast, dependency-free repository scanner",
	"• Curated rules for documentation, structure, and configuration",
	"• Automated fixes for common issues",
	"• JSON, YAML, Markdown, and pretty outputs",
	"• Regression gating for CI pipelines",
	"• GitHub Action for PR comments",
	"• SVG badge generation",
}

// var typicalWorkflowList = []string{
// 	"Scan the repository",
// 	"Run full diagnostics to identify issues and compute scores",
// 	"Export a snapshot of the current health state (before changes)",
// 	"Apply an improvement discovered from running `hewd doctor`",
// 	"Export a new snapshot after making improvements",
// 	"Compare reports using the diff engine",
// }

func buildHewdLogo(s string) string {
	return formatter.CyanBold(s)
}

func buildLogoTagline(s string) string {
	return formatter.WhiteBoldItalic(s)
}

func buildCyanText(s string) string {
	return formatter.Cyan(s)
}

func buildHeader(s string) string {
	return formatter.WhiteBold(s)
}

func buildFeaturesList(features []string) string {
	var featuresList strings.Builder
	featuresList.WriteString(`
  
`)

	const tab = `  `

	for _, feature := range features {
		featuresList.WriteString(tab + feature + "\n")
	}

	return featuresList.String()

}

func buildUseLine(start, name, end string) string {
	return start + buildCyanText(name) + end
}

func buildCyanBoldItalicText(s string) string {
	return formatter.CyanBoldItalic(s)
}

// Typical workflow
var RootTypicalWorkflowHeader = buildCyanBoldItalicText("Typical workflow:")

var ScanRepo = formatter.Blue("1. Scan the repository")
var DoctorRepo = formatter.Blue("2. Run full diagnostics to identify issues and compute scores")
var FirstExport = formatter.Blue("3. Export a snapshot of the current health state (before changes)")
var ApplyImprovement = formatter.Blue("4. Apply an improvement discovered from running `hewd doctor`")
var SecondExport = formatter.Blue("5. Export a new snapshot after making improvements")
var CompareDiff = formatter.Blue("6. Compare reports using the diff engine")

var rootLongTypicalWorkflow = RootTypicalWorkflowHeader + `

  ` + ScanRepo + `
  hewd scan

  ` + DoctorRepo + `
  hewd doctor

  ` + FirstExport + `
  hewd export --output old.json

  ` + ApplyImprovement + `
  # e.g., create a LICENSE file, add a README section, or fix a config issue

  ` + SecondExport + `
  hewd export --json --output new.json

  ` + CompareDiff + `
  hewd diff old.json new.json --md > diff.md

`

// Root header
var rootLongHeader = buildHewdLogo("hewd") + buildLogoTagline(" – repository health diagnostics, scoring, and automated fixes\n\n")

// Purpose shown below tagline
var rootLongPurpose = buildCyanText("hewd") + ` analyzes repository health by evaluating documentation, configuration,
and structural conventions. It provides fast scanning, actionable feedback,
health scores, diff reports, and automated fixes.

`

// Features shown below purpose
var rootLongFeatures = buildHeader("Features:") + buildFeaturesList(featuresList) + buildUseLine(`
Use `, "hewd", ` to maintain consistent documentation, detect regressions, enforce
standards, and track repository maturity over time.

`)

// Typical workflow

// For use in "pkg/cmd/root.go"
const RootUse = "hewd [command] [flags]"

const RootShort = "Analyze, score, and improve the health of software repositories."

// Full description.
var RootLong = rootLongHeader + rootLongPurpose + rootLongFeatures + rootLongTypicalWorkflow

var start = formatter.CyanItalic(`(See `)
var wfc = formatter.CyanItalic(`"Typical workflow"`)
var end = formatter.CyanItalic(` above for how these commands fit together — all commands shown are safe to run.)`)

var parenthesisClean = start + wfc + end

// Examples shown in `hewd --help`.
var RootExample = parenthesisClean + `

  # Scan a repository
  hewd scan

  # Run full diagnostics
  hewd doctor

  # Export machine-readable project health
  hewd export --json --output new.json

  # Compare reports using the diff engine
  hewd diff old.json new.json --md > diff.md

  # Apply automated fixes
  hewd fix --apply

  # Generate an SVG badge
  hewd badge --output badge.svg

  # Initialize a hewd configuration file
  hewd init
  `
