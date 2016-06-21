package bazaar

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".bzr/",
		".bzrignore",
		"",
	}
	template.Add("github/global/bazaar", strings.Join(ignore, "\n"))
}
