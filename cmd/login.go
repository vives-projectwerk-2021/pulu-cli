package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.AddCommand(loginProductionCmd)
	loginCmd.AddCommand(loginStagingCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login",
}

var loginProductionCmd = &cobra.Command{
	Use:   "production",
	Short: "log into the production server",
	Run: func(cmd *cobra.Command, args []string) {
		proxy.Shell(1)
	},
}
var loginStagingCmd = &cobra.Command{
	Use:   "staging",
	Short: "log into the staging server",
	Run: func(cmd *cobra.Command, args []string) {
		proxy.Shell(2)
	},
}
