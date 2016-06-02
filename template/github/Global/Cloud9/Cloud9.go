package Cloud9

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# Cloud9 IDE - http://c9.io",
		".c9revisions",
		".c9",
	}
	template.SetTemplate("GitHub/Global/Cloud9", strings.Join(ignore, "\n"))
}
