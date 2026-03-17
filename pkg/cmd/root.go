package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCmd(ver string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hewd",
		Short: "hewd — Documentation & project automation CLI",
		Long: `hewd is a command-line tool for initializing, scanning,
and managing documentation assets for codebases.`,
	}

	// Global persistent flags can be added here
	cmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")

	// Built-in version command
	cmd.AddCommand(newVersionCmd(ver))

	// Add init command
	cmd.AddCommand(newInitCmd())

	// Add scan command
	cmd.AddCommand(newScanCmd())

	// Add doctor command
	cmd.AddCommand(newDoctorCmd())

	return cmd
}

func newVersionCmd(ver string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hewd version:", ver)
		},
	}
}
