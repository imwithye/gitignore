package redcar

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".redcar",
		"",
	}
	template.Add("github/global/redcar", strings.Join(ignore, "\n"))
}
