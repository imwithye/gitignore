package cmd

import (
	"fmt"
	"github.com/imwithye/gitignore/template"
	_ "github.com/imwithye/gitignore/template/all"
	"github.com/spf13/cobra"
	"strings"
)

var showCmd = &cobra.Command{
	Use:   "show [templates]",
	Short: "Show what will be added to .gitignore",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
		ignores := []string{}
		for _, arg := range args {
			t, err := template.Get(arg)
			if err != nil {
				continue
			}
			ignores = append(ignores, t.Ignore)
		}
		if len(ignores) > 0 {
			fmt.Print(strings.Join(ignores, "\n\n"))
		}
	},
}

func init() {
	RootCmd.AddCommand(showCmd)
}
