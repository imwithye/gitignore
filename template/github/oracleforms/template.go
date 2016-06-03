package oracleforms

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Compiled Form Modules",
		"*.fmx",
		"",
		"# Compiled Menu Modules",
		"*.mmx",
		"",
		"# Compiled Pre-Linked Libraries",
		"*.plx",
		"",
	}
	template.Add("github/oracleforms", strings.Join(ignore, "\n"))
}
