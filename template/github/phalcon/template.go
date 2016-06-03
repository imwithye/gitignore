package phalcon

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"/cache/",
		"/config/development/",
		"",
	}
	template.Add("github/phalcon", strings.Join(ignore, "\n"))
}
