package mercurial

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".hg/",
		".hgignore",
		".hgsigs",
		".hgsub",
		".hgsubstate",
		".hgtags",
		"",
	}
	template.Add("github/global/mercurial", strings.Join(ignore, "\n"))
}
