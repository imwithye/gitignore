package SketchUp

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"*.skb",
	}
	template.SetTemplate("GitHub/SketchUp", strings.Join(ignore, "\n"))
}
