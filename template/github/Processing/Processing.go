package Processing

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		".DS_Store",
		"applet",
		"application.linux32",
		"application.linux64",
		"application.windows32",
		"application.windows64",
		"application.macosx",
	}
	template.SetTemplate("GitHub/Processing", strings.Join(ignore, "\n"))
}
