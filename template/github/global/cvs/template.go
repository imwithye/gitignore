package cvs

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"/CVS/*",
		"**/CVS/*",
		".cvsignore",
		"*/.cvsignore",
		"",
	}
	template.Add("github/global/cvs", strings.Join(ignore, "\n"))
}
