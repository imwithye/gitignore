package forcedotcom

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".project",
		".settings",
		"salesforce.schema",
		"Referenced Packages",
		"",
	}
	template.Add("github/forcedotcom", strings.Join(ignore, "\n"))
}
