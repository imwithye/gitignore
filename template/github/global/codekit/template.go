package codekit

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# General CodeKit files to ignore",
		"config.codekit",
		"/min",
		"",
	}
	template.Add("github/global/codekit", strings.Join(ignore, "\n"))
}
