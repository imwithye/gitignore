package Kohana

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"application/cache/*",
		"application/logs/*",
	}
	template.SetTemplate("GitHub/Kohana", strings.Join(ignore, "\n"))
}
