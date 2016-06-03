package gpg

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"secring.*",
		"",
		"",
	}
	template.Add("github/global/gpg", strings.Join(ignore, "\n"))
}
