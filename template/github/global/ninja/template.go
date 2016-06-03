package ninja

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".ninja_deps",
		".ninja_log",
		"",
	}
	template.Add("github/global/ninja", strings.Join(ignore, "\n"))
}
