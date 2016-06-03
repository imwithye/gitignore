package svn

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".svn/",
		"",
	}
	template.Add("github/global/svn", strings.Join(ignore, "\n"))
}
