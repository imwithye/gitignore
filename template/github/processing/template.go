package processing

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".DS_Store",
		"applet",
		"application.linux32",
		"application.linux64",
		"application.windows32",
		"application.windows64",
		"application.macosx",
		"",
	}
	template.Add("github/processing", strings.Join(ignore, "\n"))
}
