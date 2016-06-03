package kohana

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"application/cache/*",
		"application/logs/*",
		"",
	}
	template.Add("github/kohana", strings.Join(ignore, "\n"))
}
