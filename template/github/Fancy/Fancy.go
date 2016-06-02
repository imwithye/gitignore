package Fancy

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"*.rbc",
		"*.fyc",
	}
	template.SetTemplate("GitHub/Fancy", strings.Join(ignore, "\n"))
}
