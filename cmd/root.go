package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "pulu",
		Short: "pulu is a command line interface for managing the pulu infrastructure.",
	}
)

func Execute() error {
	return rootCmd.Execute()
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
	"-o",
	"HostKeyAlgorithms=+ssh-rsa",
}
