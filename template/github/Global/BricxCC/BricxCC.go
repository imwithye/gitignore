package BricxCC

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# Bricx Command Center IDE",
		"# http://bricxcc.sourceforge.net",
		"*.bak",
		"*.sym",
	}
	template.SetTemplate("GitHub/Global/BricxCC", strings.Join(ignore, "\n"))
}
