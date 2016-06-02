package Agda

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"*.agdai",
	}
	template.SetTemplate("GitHub/Agda", strings.Join(ignore, "\n"))
}
