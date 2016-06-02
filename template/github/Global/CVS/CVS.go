package CVS

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"/CVS/*",
		"**/CVS/*",
		".cvsignore",
		"*/.cvsignore",
	}
	template.SetTemplate("GitHub/Global/CVS", strings.Join(ignore, "\n"))
}
