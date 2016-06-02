package DartEditor

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		".project",
		".buildlog",
	}
	template.SetTemplate("GitHub/Global/DartEditor", strings.Join(ignore, "\n"))
}
