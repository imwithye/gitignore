package Idris

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"*.ibc",
		"*.o",
	}
	template.SetTemplate("GitHub/Idris", strings.Join(ignore, "\n"))
}
