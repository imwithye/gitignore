package Ninja

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		".ninja_deps",
		".ninja_log",
	}
	template.SetTemplate("GitHub/Global/Ninja", strings.Join(ignore, "\n"))
}
