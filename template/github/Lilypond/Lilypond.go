package Lilypond

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"*.pdf",
		"*.ps",
		"*.midi",
		"*.mid",
		"*.log",
		"*~",
	}
	template.SetTemplate("GitHub/Lilypond", strings.Join(ignore, "\n"))
}
