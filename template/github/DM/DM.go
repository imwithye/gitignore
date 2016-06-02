package DM

import "strings"
import "github.com/imwithye/git_ignore/template"

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
