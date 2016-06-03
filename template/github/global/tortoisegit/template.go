package tortoisegit

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Project-level settings",
		"/.tgitconfig",
		"",
	}
	template.Add("github/global/tortoisegit", strings.Join(ignore, "\n"))
}
