package AppEngine

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# Google App Engine generated folder",
		"appengine-generated/",
	}
	template.SetTemplate("GitHub/AppEngine", strings.Join(ignore, "\n"))
}
