package momentics

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/global/momentics", strings.Join(ignore, "\n"))
}
