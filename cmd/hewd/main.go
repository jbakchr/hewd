package main

import (
	"fmt"
	"os"

	"github.com/jbakchr/hewd/internal/cliutils"
	"github.com/jbakchr/hewd/internal/version"
	"github.com/jbakchr/hewd/pkg/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd(version.Version)

	if err := rootCmd.Execute(); err != nil {
		// Determine executed command based on os.Args:
		// hewd <command> [...]
		cmdName := determineExecutedCommand(os.Args)

		// If this is one of our structured HewdError values:
		if he, ok := err.(cliutils.HewdError); ok {
			// Top‑level formatting:
			fmt.Fprintln(os.Stderr, cliutils.RootErr(cmdName, he.Msg, he.Hint))
		} else {
			// Fallback for unexpected error types:
			fmt.Fprintf(os.Stderr,
				"%serror (hewd %s):%s %v\n",
				cliutils.Red, cmdName, cliutils.Reset,
				err,
			)
		}

		os.Exit(1)
	}
}

// Returns the name of the executed subcommand.
// Example: "doctor" in:  hewd doctor --json
func determineExecutedCommand(args []string) string {
	if len(args) <= 1 {
		return "hewd"
	}
	return args[1]
}
