package cmd

import (
	"fmt"

	"github.com/jbakchr/hewd/internal/helptext"
	"github.com/spf13/cobra"
)

func NewRootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     helptext.RootUse,
		Short:   helptext.RootShort,
		Long:    helptext.RootLong,
		Example: helptext.RootExample,
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
