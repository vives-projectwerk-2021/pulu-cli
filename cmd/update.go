package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(updateProductionCmd)
	updateCmd.AddCommand(updateStagingCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the pulu server",
}

var updateProductionCmd = &cobra.Command{
	Use:   "production",
	Short: "Update the production server",
	Run: func(cmd *cobra.Command, args []string) {
		execProductionCmd.Run(cmd, []string{"make -C deployment production-update production"})
	},
}

var updateStagingCmd = &cobra.Command{
	Use:   "staging",
	Short: "Update the staging server",
	Run: func(cmd *cobra.Command, args []string) {
		execStagingCmd.Run(cmd, []string{"make -C deployment staging-update staging"})
	},
}
