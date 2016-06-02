package Mercurial

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		".hg/",
		".hgignore",
		".hgsigs",
		".hgsub",
		".hgsubstate",
		".hgtags",
	}
	template.SetTemplate("GitHub/Global/Mercurial", strings.Join(ignore, "\n"))
}
