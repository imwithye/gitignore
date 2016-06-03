package netbeans

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"nbproject/private/",
		"build/",
		"nbbuild/",
		"dist/",
		"nbdist/",
		"nbactions.xml",
		".nb-gradle/",
		"",
	}
	template.Add("github/global/netbeans", strings.Join(ignore, "\n"))
}
