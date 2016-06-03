package idris

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.ibc",
		"*.o",
		"",
	}
	template.Add("github/idris", strings.Join(ignore, "\n"))
}
