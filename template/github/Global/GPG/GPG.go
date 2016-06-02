package GPG

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"secring.*",
		"",
	}
	template.SetTemplate("GitHub/Global/GPG", strings.Join(ignore, "\n"))
}
