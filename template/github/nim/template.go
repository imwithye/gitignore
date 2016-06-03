package nim

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"nimcache/",
		"",
	}
	template.Add("github/nim", strings.Join(ignore, "\n"))
}
