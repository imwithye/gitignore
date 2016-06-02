package Redcar

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		".redcar",
	}
	template.SetTemplate("GitHub/Global/Redcar", strings.Join(ignore, "\n"))
}
