package Scrivener

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"/Files/binder.autosave",
		"/Files/binder.backup",
		"/Files/search.indexes",
		"/Files/user.lock",
		"/Files/Docs/docs.checksum",
		"/QuickLook/",
		"/Settings/ui.plist",
	}
	template.SetTemplate("GitHub/Scrivener", strings.Join(ignore, "\n"))
}
