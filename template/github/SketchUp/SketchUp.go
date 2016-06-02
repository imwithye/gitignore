package SketchUp

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"*.skb",
	}
	template.SetTemplate("GitHub/SketchUp", strings.Join(ignore, "\n"))
}
