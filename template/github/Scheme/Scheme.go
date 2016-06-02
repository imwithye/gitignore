package Scheme

import "strings"
import "github.com/imwithye/git_ignore/template"

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
