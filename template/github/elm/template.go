package elm

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# elm-package generated files",
		"elm-stuff/",
		"# elm-repl generated files",
		"repl-temp-*",
		"",
	}
	template.Add("github/elm", strings.Join(ignore, "\n"))
}
