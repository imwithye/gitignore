package OracleForms

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Compiled Form Modules",
		"*.fmx",
		"",
		"# Compiled Menu Modules",
		"*.mmx",
		"",
		"# Compiled Pre-Linked Libraries",
		"*.plx",
	}
	template.SetTemplate("GitHub/OracleForms", strings.Join(ignore, "\n"))
}
