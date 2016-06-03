package yeoman

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"node_modules/",
		"bower_components/",
		"*.log",
		"",
		"build/",
		"dist/",
		"",
	}
	template.Add("github/yeoman", strings.Join(ignore, "\n"))
}
