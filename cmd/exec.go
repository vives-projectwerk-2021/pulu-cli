package cmd

import (
	"os"
	"os/exec"

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
		baseSSHargs = append(baseSSHargs, "exec", "production")
		baseSSHargs = append(baseSSHargs, args...)
		exeCmd := exec.Command("ssh", baseSSHargs...)
		exeCmd.Stdout = os.Stdout
		exeCmd.Stdin = os.Stdin
		exeCmd.Stderr = os.Stderr
		exeCmd.Run()
	},
}
var execStagingCmd = &cobra.Command{
	Use:   "staging",
	Short: "execute a command om the staging server",
	Run: func(cmd *cobra.Command, args []string) {
		baseSSHargs = append(baseSSHargs, "exec", "staging")
		baseSSHargs = append(baseSSHargs, args...)
		exeCmd := exec.Command("ssh", baseSSHargs...)
		exeCmd.Stdout = os.Stdout
		exeCmd.Stdin = os.Stdin
		exeCmd.Stderr = os.Stderr
		exeCmd.Run()
	},
}
