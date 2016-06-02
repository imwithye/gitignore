package EiffelStudio

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# The compilation directory",
		"EIFGENs",
	}
	template.SetTemplate("GitHub/Global/EiffelStudio", strings.Join(ignore, "\n"))
}
