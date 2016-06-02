package DM

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"*.dmb",
		"*.rsc",
		"*.int",
		"*.lk",
		"gitignore-master.zip",
	}
	template.SetTemplate("GitHub/DM", strings.Join(ignore, "\n"))
}
