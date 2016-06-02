package NetBeans

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"nbproject/private/",
		"build/",
		"nbbuild/",
		"dist/",
		"nbdist/",
		"nbactions.xml",
		".nb-gradle/",
	}
	template.SetTemplate("GitHub/Global/NetBeans", strings.Join(ignore, "\n"))
}
