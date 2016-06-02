package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.1.0")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
