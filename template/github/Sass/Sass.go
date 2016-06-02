package Sass

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		".sass-cache/",
		"*.css.map",
	}
	template.SetTemplate("GitHub/Sass", strings.Join(ignore, "\n"))
}
