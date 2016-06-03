package fancy

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.rbc",
		"*.fyc",
		"",
	}
	template.Add("github/fancy", strings.Join(ignore, "\n"))
}
