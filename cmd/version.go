package cmd

import (
	"fmt"
	"github.com/imwithye/gitignore/cmd/platform"
	"github.com/imwithye/gitignore/cmd/version"
	"github.com/spf13/cobra"
)

var versionCmdNoCheck bool
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s, %s, %s.\n", version.CurrentVersion, platform.Platform, "Ciel <me@ciel.im>")
		if versionCmdNoCheck {
			return
		}
		latestVersion, url, err := version.Latest()
		if err != nil || version.CurrentVersion.Cmp(latestVersion) > 0 {
			fmt.Println("Already up-to-date.")
			return
		}
		fmt.Printf("Latest version %s, %s\n", latestVersion, url)
	},
}

func init() {
	versionCmd.Flags().BoolVarP(&versionCmdNoCheck, "no-check", "", false, "don't check latest version")
	RootCmd.AddCommand(versionCmd)
}
