package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.3.0, Ciel <me@ciel.im>")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
