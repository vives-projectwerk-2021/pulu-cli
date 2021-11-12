package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(execCmd)
	execCmd.AddCommand(execProductionCmd)
	execCmd.AddCommand(execStagingCmd)
}

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "execute a command on the server",
}

var execProductionCmd = &cobra.Command{
	Use:   "production",
	Short: "execute a command om the production server",
	Run: func(cmd *cobra.Command, args []string) {
		proxy.Execute(1, args...)
	},
}
var execStagingCmd = &cobra.Command{
	Use:   "staging",
	Short: "execute a command om the staging server",
	Run: func(cmd *cobra.Command, args []string) {
		proxy.Execute(2, args...)
	},
}
