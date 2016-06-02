package VVVV

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"",
		"# .v4p backup files",
		"*~.xml",
		"",
		"# Dynamic plugins .dll",
		"bin/",
	}
	template.SetTemplate("GitHub/VVVV", strings.Join(ignore, "\n"))
}
