package vvvv

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"",
		"# .v4p backup files",
		"*~.xml",
		"",
		"# Dynamic plugins .dll",
		"bin/",
		"",
	}
	template.Add("github/vvvv", strings.Join(ignore, "\n"))
}
