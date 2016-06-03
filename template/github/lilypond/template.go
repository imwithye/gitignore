package lilypond

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.pdf",
		"*.ps",
		"*.midi",
		"*.mid",
		"*.log",
		"*~",
		"",
	}
	template.Add("github/lilypond", strings.Join(ignore, "\n"))
}
