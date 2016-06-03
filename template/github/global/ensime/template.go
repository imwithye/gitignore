package ensime

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Ensime specific",
		".ensime",
		".ensime_cache/",
		".ensime_lucene/",
		"",
	}
	template.Add("github/global/ensime", strings.Join(ignore, "\n"))
}
