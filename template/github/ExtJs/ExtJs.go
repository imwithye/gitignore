package ExtJs

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		".architect",
		"bootstrap.json",
		"build/",
		"ext/",
	}
	template.SetTemplate("GitHub/ExtJs", strings.Join(ignore, "\n"))
}
