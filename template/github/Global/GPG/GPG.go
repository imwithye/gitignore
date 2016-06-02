package GPG

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"secring.*",
		"",
	}
	template.SetTemplate("GitHub/Global/GPG", strings.Join(ignore, "\n"))
}
