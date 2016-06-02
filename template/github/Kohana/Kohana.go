package Kohana

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"application/cache/*",
		"application/logs/*",
	}
	template.SetTemplate("GitHub/Kohana", strings.Join(ignore, "\n"))
}
