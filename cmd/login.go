package cmd

import (
	"os"
	"os/exec"

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
		baseSSHargs = append(baseSSHargs, "login", "production")
		exeCmd := exec.Command("ssh", baseSSHargs...)
		exeCmd.Stdout = os.Stdout
		exeCmd.Stdin = os.Stdin
		exeCmd.Stderr = os.Stderr
		exeCmd.Run()
	},
}
var loginStagingCmd = &cobra.Command{
	Use:   "staging",
	Short: "log into the staging server",
	Run: func(cmd *cobra.Command, args []string) {
		baseSSHargs = append(baseSSHargs, "login", "staging")
		exeCmd := exec.Command("ssh", baseSSHargs...)
		exeCmd.Stdout = os.Stdout
		exeCmd.Stdin = os.Stdin
		exeCmd.Stderr = os.Stderr
		exeCmd.Run()
	},
}
