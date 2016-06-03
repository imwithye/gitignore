package agda

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.agdai",
		"",
	}
	template.Add("github/agda", strings.Join(ignore, "\n"))
}
