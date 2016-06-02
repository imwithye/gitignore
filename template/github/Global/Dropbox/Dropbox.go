package Dropbox

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# Dropbox settings and caches",
		".dropbox",
		".dropbox.attr",
		".dropbox.cache",
	}
	template.SetTemplate("GitHub/Global/Dropbox", strings.Join(ignore, "\n"))
}
