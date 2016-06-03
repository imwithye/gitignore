package eiffelstudio

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# The compilation directory",
		"EIFGENs",
		"",
	}
	template.Add("github/global/eiffelstudio", strings.Join(ignore, "\n"))
}
