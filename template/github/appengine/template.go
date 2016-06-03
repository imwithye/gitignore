package appengine

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Google App Engine generated folder",
		"appengine-generated/",
		"",
	}
	template.Add("github/appengine", strings.Join(ignore, "\n"))
}
