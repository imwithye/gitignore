package Espresso

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"*.esproj",
	}
	template.SetTemplate("GitHub/Global/Espresso", strings.Join(ignore, "\n"))
}
