package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "pulu",
		Short: "pulu is a command line interface for managing the pulu infrastructure.",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
