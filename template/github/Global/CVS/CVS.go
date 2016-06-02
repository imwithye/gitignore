package CVS

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"/CVS/*",
		"**/CVS/*",
		".cvsignore",
		"*/.cvsignore",
	}
	template.SetTemplate("GitHub/Global/CVS", strings.Join(ignore, "\n"))
}
