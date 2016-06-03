package dm

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.dmb",
		"*.rsc",
		"*.int",
		"*.lk",
		"*.zip",
		"",
	}
	template.Add("github/dm", strings.Join(ignore, "\n"))
}
