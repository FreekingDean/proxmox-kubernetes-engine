package main

import (
	"fmt"
	"os"
)

func main() {
	if err := NewModuleCommand(
		"pke",
		"Proxmox Kubernetes Engine CLI",
	).Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
