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

var baseSSHargs = []string{
	"-tt",
	"devops@pulu.trikthom.com",
	"-p",
	"2222",
	"-o",
	"StrictHostKeyChecking=no",
	"-o",
	"UserKnownHostsFile=/dev/null",
	"-o",
	"LogLevel=QUIET",
}

var loginProductionCmd = &cobra.Command{
	Use:   "production",
	Short: "log into the production server",
	Run: func(cmd *cobra.Command, args []string) {
		baseSSHargs = append(baseSSHargs, "production")
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
		baseSSHargs = append(baseSSHargs, "staging")
		exeCmd := exec.Command("ssh", baseSSHargs...)
		exeCmd.Stdout = os.Stdout
		exeCmd.Stdin = os.Stdin
		exeCmd.Stderr = os.Stderr
		exeCmd.Run()
	},
}
