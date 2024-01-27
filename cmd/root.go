package cmd

import (
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "",
		Short: "",
	}
	rootCmd.AddCommand(
		bgsCmd(),
		serverCmd(),
	)
	return rootCmd
}
