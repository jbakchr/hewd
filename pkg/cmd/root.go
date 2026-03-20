package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "hewd",
		Short: "hewd is a project health, documentation, and structure toolkit.",
		Long: `hewd is a fast, dependency-free toolkit for analyzing and improving the
health of software repositories. It evaluates documentation, configuration, and 
structural conventions using a curated rule engine and provides clear, actionable 
feedback, health scores, and automated fixes.

hewd supports multiple output formats (pretty, JSON, YAML, Markdown), a 
machine-readable export schema, a powerful diff engine, regression gating for 
CI workflows, and a first-class GitHub Action that posts or updates pull 
request comments.

Use hewd to maintain consistent documentation, enforce structure across 
repositories, detect regressions, and track project maturity over time.`,
		Example: `
  # Scan the repository and show a high-level summary
  hewd scan --pretty

  # Run full diagnostics and generate a Markdown report
  hewd doctor --md > health.md

  # Export machine-readable project health
  hewd export --output hewd.json

  # Compare two reports to detect new/resolved issues
  hewd diff old.json new.json --md > diff.md

  # Automatically fix missing documentation or CI files
  hewd fix --apply

  # Generate an SVG badge showing project health score
  hewd badge --output badge.svg

  # Initialize a .hewd/config.yaml configuration file
  hewd init
`,
	}

	rootCmd.AddCommand(newScanCmd())
	rootCmd.AddCommand(newDoctorCmd())
	rootCmd.AddCommand(newInitCmd())
	rootCmd.AddCommand(newBadgeCmd())
	rootCmd.AddCommand(newFixCmd())
	rootCmd.AddCommand(newExportCmd())
	rootCmd.AddCommand(newDiffCmd())

	rootCmd.Version = version
	rootCmd.SetVersionTemplate("hewd version {{.Version}}\n")

	return rootCmd
}

func Execute(version string) {
	if err := NewRootCmd(version).Execute(); err != nil {
		fmt.Println(err)
	}
}
