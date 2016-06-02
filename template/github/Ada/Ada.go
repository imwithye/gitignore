package Ada

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# Object file",
		"*.o",
		"",
		"# Ada Library Information",
		"*.ali",
	}
	template.SetTemplate("GitHub/Ada", strings.Join(ignore, "\n"))
}
