package sass

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".sass-cache/",
		"*.css.map",
		"",
	}
	template.Add("github/sass", strings.Join(ignore, "\n"))
}
