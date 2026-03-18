package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "hewd",
		Short: "hewd — Documentation & Project Structure Toolkit",
		Long:  `hewd scans your project and runs structural & documentation rules.`,
	}

	rootCmd.AddCommand(newScanCmd())
	rootCmd.AddCommand(newDoctorCmd())
	rootCmd.AddCommand(newInitCmd())
	rootCmd.AddCommand(newBadgeCmd())

	rootCmd.Version = version
	rootCmd.SetVersionTemplate("hewd version {{.Version}}\n")

	return rootCmd
}

func Execute(version string) {
	if err := NewRootCmd(version).Execute(); err != nil {
		fmt.Println(err)
	}
}
