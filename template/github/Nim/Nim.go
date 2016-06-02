package Nim

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"nimcache/",
	}
	template.SetTemplate("GitHub/Nim", strings.Join(ignore, "\n"))
}
