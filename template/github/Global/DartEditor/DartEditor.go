package DartEditor

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		".project",
		".buildlog",
	}
	template.SetTemplate("GitHub/Global/DartEditor", strings.Join(ignore, "\n"))
}
