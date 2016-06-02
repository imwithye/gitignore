package Waf

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# for projects that use Waf for building: http://code.google.com/p/waf/",
		".waf-*",
		".waf3-*",
		".lock-*",
	}
	template.SetTemplate("GitHub/Waf", strings.Join(ignore, "\n"))
}
