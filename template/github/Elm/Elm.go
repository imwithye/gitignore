package Elm

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# elm-package generated files",
		"elm-stuff/",
		"# elm-repl generated files",
		"repl-temp-*",
	}
	template.SetTemplate("GitHub/Elm", strings.Join(ignore, "\n"))
}
