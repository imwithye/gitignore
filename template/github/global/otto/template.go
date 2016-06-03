package otto

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".otto/",
		"",
	}
	template.Add("github/global/otto", strings.Join(ignore, "\n"))
}
