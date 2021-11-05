package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(proxyCmd)
	proxyCmd.AddCommand(proxyProductionCmd)
	proxyCmd.AddCommand(proxyStagingCmd)
}

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Creates a proxy server or application-level gateway between localhost and the Pulu Server.",
}

var proxyProductionCmd = &cobra.Command{
	Use: "production",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			exeCmd := exec.Command("ssh", append(baseSSHargs, "proxy", "production", args[0])...)
			exeCmd.Stdout = os.Stdout
			exeCmd.Stdin = os.Stdin
			exeCmd.Stderr = os.Stderr
			exeCmd.Run()
		}
	},
}

var proxyStagingCmd = &cobra.Command{
	Use: "staging",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			exeCmd := exec.Command("ssh", append(baseSSHargs, "proxy", "staging", args[0])...)
			exeCmd.Stdout = os.Stdout
			exeCmd.Stdin = os.Stdin
			exeCmd.Stderr = os.Stderr
			exeCmd.Run()
		}
	},
}
