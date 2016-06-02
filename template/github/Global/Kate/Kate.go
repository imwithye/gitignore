package Kate

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# Swap Files #",
		".*.kate-swp",
		".swp.*",
	}
	template.SetTemplate("GitHub/Global/Kate", strings.Join(ignore, "\n"))
}
