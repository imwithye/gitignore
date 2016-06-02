package ExtJs

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		".architect",
		"bootstrap.json",
		"build/",
		"ext/",
	}
	template.SetTemplate("GitHub/ExtJs", strings.Join(ignore, "\n"))
}
