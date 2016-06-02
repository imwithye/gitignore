package SVN

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		".svn/",
	}
	template.SetTemplate("GitHub/Global/SVN", strings.Join(ignore, "\n"))
}
