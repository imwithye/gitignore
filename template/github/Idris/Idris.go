package Idris

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"*.ibc",
		"*.o",
	}
	template.SetTemplate("GitHub/Idris", strings.Join(ignore, "\n"))
}
