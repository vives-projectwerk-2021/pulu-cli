package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

func init() {
	// rootCmd.AddCommand(proxyCmd)
	// proxyCmd.AddCommand(proxyProductionCmd)
	// proxyCmd.AddCommand(proxyStagingCmd)
}

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Creates a proxy server or application-level gateway between localhost and the Pulu Server.",
	Run: func(cmd *cobra.Command, args []string) {
		proxyStagingCmd.Run(cmd, args)
	},
}

var proxyProductionCmd = &cobra.Command{
	Use: "production",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("prod")
	},
}

var proxyStagingCmd = &cobra.Command{
	Use: "staging",
	Run: func(cmd *cobra.Command, args []string) {
		baseSSHargs = append(baseSSHargs, "proxy staging")
		exeCmd := exec.Command("ssh", baseSSHargs...)
		exeCmd.Start()

		sigs := make(chan os.Signal, 1)
		done := make(chan bool, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			sig := <-sigs
			fmt.Println()
			fmt.Println(sig)
			done <- true
		}()

		fmt.Println("awaiting signal")
		<-done
		fmt.Println("exiting")
		exeCmd.Process.Signal(os.Kill)
		// exeCmd.Process.Kill()
	},
}
