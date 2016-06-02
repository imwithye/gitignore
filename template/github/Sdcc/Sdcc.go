package Sdcc

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# SDCC stuff",
		"*.lnk",
		"*.lst",
		"*.map",
		"*.mem",
		"*.rel",
		"*.rst",
		"*.sym",
	}
	template.SetTemplate("GitHub/Sdcc", strings.Join(ignore, "\n"))
}
