package scrivener

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"/Files/binder.autosave",
		"/Files/binder.backup",
		"/Files/search.indexes",
		"/Files/user.lock",
		"/Files/Docs/docs.checksum",
		"/QuickLook/",
		"/Settings/ui.plist",
		"",
	}
	template.Add("github/scrivener", strings.Join(ignore, "\n"))
}
