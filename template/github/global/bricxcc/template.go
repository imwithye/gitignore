package bricxcc

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Bricx Command Center IDE",
		"# http://bricxcc.sourceforge.net",
		"*.bak",
		"*.sym",
		"",
	}
	template.Add("github/global/bricxcc", strings.Join(ignore, "\n"))
}
