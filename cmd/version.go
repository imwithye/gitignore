package cmd

import (
	"fmt"
	"github.com/imwithye/gitignore/cmd/platform"
	"github.com/imwithye/gitignore/cmd/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s, %s, %s\n", version.CurrentVersion, platform.Platform, "Ciel <me@ciel.im>")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
