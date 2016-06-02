package Phalcon

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"/cache/",
		"/config/development/",
	}
	template.SetTemplate("GitHub/Phalcon", strings.Join(ignore, "\n"))
}
