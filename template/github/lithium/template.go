package lithium

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"libraries/*",
		"resources/tmp/*",
		"",
	}
	template.Add("github/lithium", strings.Join(ignore, "\n"))
}
