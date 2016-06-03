package sdcc

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/sdcc", strings.Join(ignore, "\n"))
}
