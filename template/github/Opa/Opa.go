package Opa

import "strings"
import "github.com/imwithye/git_ignore/template"

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
	}
	template.SetTemplate("GitHub/Opa", strings.Join(ignore, "\n"))
}
