package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login [production/staging]",
	Short: "login",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		a := []string{
			"-tt",
			"devops@pulu.trikthom.com",
			"-p",
			"2222",
			"-o",
			"StrictHostKeyChecking=no",
		}

		a = append(a, args...)

		if !strings.EqualFold(args[0], "production") && !strings.EqualFold(args[0], "staging") {
			fmt.Println("Use production or staging")
			return
		}

		exeCmd := exec.Command("ssh", a...)

		exeCmd.Stdout = os.Stdout
		exeCmd.Stdin = os.Stdin
		exeCmd.Stderr = os.Stderr
		exeCmd.Run()
	},
}
