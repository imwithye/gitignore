package scheme

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.ss~",
		"*.ss#*",
		".#*.ss",
		"",
		"*.scm~",
		"*.scm#*",
		".#*.scm",
		"",
	}
	template.Add("github/scheme", strings.Join(ignore, "\n"))
}
