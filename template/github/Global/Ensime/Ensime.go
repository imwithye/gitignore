package Ensime

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Ensime specific",
		".ensime",
		".ensime_cache/",
		".ensime_lucene/",
	}
	template.SetTemplate("GitHub/Global/Ensime", strings.Join(ignore, "\n"))
}
