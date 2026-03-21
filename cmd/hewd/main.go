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

		// Determine executed command from args
		cmdName := determineExecutedCommand(os.Args)

		// If it's a HewdError, format with RootErr
		if he, ok := err.(cliutils.HewdError); ok {
			fmt.Fprintln(os.Stderr, cliutils.RootErr(cmdName, he.Msg, he.Hint))
		} else {
			// fallback for unexpected errors
			fmt.Fprintf(os.Stderr, "%serror (hewd %s):%s %v\n",
				cliutils.Red, cmdName, cliutils.Reset, err)
		}

		os.Exit(1)
	}
}

func determineExecutedCommand(args []string) string {
	if len(args) <= 1 {
		return "hewd"
	}
	return args[1] // e.g. doctor, scan, diff, etc.
}
