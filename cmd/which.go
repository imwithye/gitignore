package cmd

import (
	"fmt"
	"github.com/imwithye/gitignore/template"
	_ "github.com/imwithye/gitignore/template/all"
	"github.com/spf13/cobra"
)

var whichCmd = &cobra.Command{
	Use:   "which",
	Short: "Show which ignore templates will be added",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
		for _, arg := range args {
			t, err := template.GetTemplate(arg)
			if err != nil {
				fmt.Println("*", err)
				continue
			}
			fmt.Println("-", t.Path)
		}
	},
}

func init() {
	RootCmd.AddCommand(whichCmd)
}
