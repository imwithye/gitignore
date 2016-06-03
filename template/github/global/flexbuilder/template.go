package flexbuilder

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"bin/",
		"bin-debug/",
		"bin-release/",
		"",
	}
	template.Add("github/global/flexbuilder", strings.Join(ignore, "\n"))
}
