package cmd

import (
	"fmt"
	"github.com/imwithye/gitignore/cmd/platform"
	"github.com/spf13/cobra"
)

const Version = "v1.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s, %s, %s\n", Version, platform.Platform, "Ciel <me@ciel.im>")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
