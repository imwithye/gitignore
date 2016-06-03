package extjs

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".architect",
		"bootstrap.json",
		"build/",
		"ext/",
		"",
	}
	template.Add("github/extjs", strings.Join(ignore, "\n"))
}
