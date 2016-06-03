package slickedit

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# SlickEdit workspace and project files are ignored by default because",
		"# typically they are considered to be developer-specific and not part of a",
		"# project.",
		"*.vpw",
		"*.vpj",
		"",
		"# SlickEdit workspace history and tag files always contain user-specific",
		"# data so they should not be stored in a repository.",
		"*.vpwhistu",
		"*.vpwhist",
		"*.vtg",
		"",
	}
	template.Add("github/global/slickedit", strings.Join(ignore, "\n"))
}
