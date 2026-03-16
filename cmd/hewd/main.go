package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jbakchr/hewd/internal/version"
)

func main() {
	// Flags
	showVersion := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println("hewd version:", version.Version)
		return
	}

	// Default behavior
	fmt.Println("hewd CLI — nothing here yet!")
	os.Exit(0)
}
