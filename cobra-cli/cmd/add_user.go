package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var AddUser = &cobra.Command{
	Use:   "add-user USERNAME",
	Short: "Add a new user",
	Long:  `add-user adds a new user to the system. An argument of USERNAME is required`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		username := strings.Join(args, " ")

		if Verbose {
			fmt.Println("[Verbose] Creating a new user:", username)
		}

		if DryRun {
			fmt.Println("[Dry-Run] Added user:", username)
		} else {
			fmt.Println("[+] Added user:", username)
		}
	},
}
