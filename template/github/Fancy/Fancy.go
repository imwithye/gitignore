package Fancy

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"*.rbc",
		"*.fyc",
	}
	template.SetTemplate("GitHub/Fancy", strings.Join(ignore, "\n"))
}
