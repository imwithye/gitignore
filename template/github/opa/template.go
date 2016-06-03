package opa

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"_build",
		"_tracks",
		"",
		"opa-debug-js",
		"",
		"*.opp",
		"*.opx",
		"*.opx.broken",
		"*.dump",
		"*.api",
		"*.api-txt",
		"*.exe",
		"*.log",
		"",
	}
	template.Add("github/opa", strings.Join(ignore, "\n"))
}
