package dropbox

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Dropbox settings and caches",
		".dropbox",
		".dropbox.attr",
		".dropbox.cache",
		"",
	}
	template.Add("github/global/dropbox", strings.Join(ignore, "\n"))
}
