package waf

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# for projects that use Waf for building: http://code.google.com/p/waf/",
		".waf-*",
		".waf3-*",
		".lock-*",
		"",
	}
	template.Add("github/waf", strings.Join(ignore, "\n"))
}
