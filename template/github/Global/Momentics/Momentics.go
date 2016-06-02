package Momentics

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Built files",
		"x86/",
		"arm/",
		"arm-p/",
		"translations/*.qm",
		"",
		"# IDE settings",
		".settings/",
	}
	template.SetTemplate("GitHub/Global/Momentics", strings.Join(ignore, "\n"))
}
