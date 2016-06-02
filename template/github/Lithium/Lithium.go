package Lithium

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"libraries/*",
		"resources/tmp/*",
	}
	template.SetTemplate("GitHub/Lithium", strings.Join(ignore, "\n"))
}
