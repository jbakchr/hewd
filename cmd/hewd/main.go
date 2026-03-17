package main

import (
	"fmt"
	"os"

	"github.com/jbakchr/hewd/internal/version"
	"github.com/jbakchr/hewd/pkg/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd(version.Version)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
