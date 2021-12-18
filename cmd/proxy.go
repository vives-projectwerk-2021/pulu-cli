package cmd

import (
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
			proxy.Proxy(1, args[0])
		}
	},
}

var proxyStagingCmd = &cobra.Command{
	Use: "staging",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			proxy.Proxy(2, args[0])
		}
	},
}
