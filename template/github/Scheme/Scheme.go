package Scheme

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"*.ss~",
		"*.ss#*",
		".#*.ss",
		"",
		"*.scm~",
		"*.scm#*",
		".#*.scm",
	}
	template.SetTemplate("GitHub/Scheme", strings.Join(ignore, "\n"))
}
