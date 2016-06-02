package Otto

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		".otto/",
	}
	template.SetTemplate("GitHub/Global/Otto", strings.Join(ignore, "\n"))
}
