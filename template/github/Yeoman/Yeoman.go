package Yeoman

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"node_modules/",
		"bower_components/",
		"*.log",
		"",
		"build/",
		"dist/",
	}
	template.SetTemplate("GitHub/Yeoman", strings.Join(ignore, "\n"))
}
