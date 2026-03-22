package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jbakchr/hewd/internal/helptext"
)

func NewRootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     helptext.RootUse,
		Short:   helptext.RootShort,
		Long:    helptext.RootLong,
		Example: helptext.RootExample,
	}

	// -------------------------
	// Define command groups
	// -------------------------
	rootCmd.AddGroup(&cobra.Group{
		ID:    "analysis",
		Title: "Analysis Commands:",
	})

	rootCmd.AddGroup(&cobra.Group{
		ID:    "maintenance",
		Title: "Maintenance Commands:",
	})

	rootCmd.AddGroup(&cobra.Group{
		ID:    "reporting",
		Title: "Reporting Commands:",
	})

	// -------------------------
	// Add commands
	// -------------------------
	rootCmd.AddCommand(newScanCmd())
	rootCmd.AddCommand(newDoctorCmd())
	rootCmd.AddCommand(newDiffCmd())

	rootCmd.AddCommand(newFixCmd())
	rootCmd.AddCommand(newInitCmd())

	rootCmd.AddCommand(newExportCmd())
	rootCmd.AddCommand(newBadgeCmd())

	rootCmd.Version = version
	rootCmd.SetVersionTemplate("hewd version {{.Version}}\n")

	return rootCmd
}
