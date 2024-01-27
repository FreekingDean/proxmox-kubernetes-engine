package main

import (
	"fmt"
	"os"
)

func main() {
	if err := NewModuleCommand(
		"pke",
		"Example CLI tool",
	).Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
