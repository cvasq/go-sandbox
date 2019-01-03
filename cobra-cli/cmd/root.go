package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DryRun bool
var Verbose bool

func init() {
	RootCmd.PersistentFlags().BoolVarP(&DryRun, "dryrun", "d", false, "Dry run")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Verbose output")

	RootCmd.AddCommand(AddUser)

}

var RootCmd = &cobra.Command{
	Use: "cobra-cli",
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			fmt.Println("Running in Verbose mode...")
		}
		fmt.Printf("testctl Version 1.0\n")
	},
}
