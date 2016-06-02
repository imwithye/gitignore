package Elisp

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# Compiled",
		"*.elc",
		"",
		"# Packaging",
		".cask",
	}
	template.SetTemplate("GitHub/Elisp", strings.Join(ignore, "\n"))
}
