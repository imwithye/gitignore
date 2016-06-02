package Phalcon

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"/cache/",
		"/config/development/",
	}
	template.SetTemplate("GitHub/Phalcon", strings.Join(ignore, "\n"))
}
