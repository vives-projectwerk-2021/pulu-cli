package cmd

import (
	"github.com/spf13/cobra"
	proxytunnel "thomascrmbz.com/proxytunnel/client"
)

var (
	rootCmd = &cobra.Command{
		Use:   "pulu",
		Short: "pulu is a command line interface for managing the pulu infrastructure.",
	}

	proxy = proxytunnel.NewProxyClient("pulu.trikthom.com", 2222)
)

func Execute() error {
	return rootCmd.Execute()
}
