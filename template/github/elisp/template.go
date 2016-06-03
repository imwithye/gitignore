package elisp

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Compiled",
		"*.elc",
		"",
		"# Packaging",
		".cask",
		"",
	}
	template.Add("github/elisp", strings.Join(ignore, "\n"))
}
